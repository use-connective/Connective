import { fetchProvider } from "./api.js";
import { BACKEND_BASE_URL } from "./index.js";
import { getCSS } from "./popup.css.js";
import { getHTML, getOverviewHTML } from "./popup.html.js";
export class IntegrationPopUpCard {
    projectId;
    userId;
    projectSecret;
    containerId = "connective-dom-popup-root";
    styleId = "connective-dom-popup-styles";
    constructor(projectId, userId, projectSecret) {
        this.projectId = projectId;
        this.userId = userId;
        this.projectSecret = projectSecret;
    }
    async open({ provider, onSuccess, onClose }) {
        if (document.getElementById(this.containerId))
            return;
        let providerData = await fetchProvider(provider);
        // Disable scrolling when connect portal is opened.
        const prevOverflow = document.body.style.overflow;
        document.body.style.overflow = "hidden";
        // Injecting CSS by creating a new style tag.
        if (!document.getElementById(this.styleId)) {
            const style = document.createElement("style");
            style.id = this.styleId;
            style.innerHTML = getCSS(providerData);
            document.head.appendChild(style);
        }
        // backdrop
        const backdrop = document.createElement("div");
        backdrop.className = "connective-backdrop";
        backdrop.setAttribute("data-connective", "backdrop");
        // popup container
        const container = document.createElement("div");
        container.id = this.containerId;
        container.className = "connective-popup-container";
        container.setAttribute("role", "dialog");
        container.setAttribute("aria-modal", "true");
        container.innerHTML = await getHTML(providerData);
        // Appending backdrop and popup container to body.
        document.body.appendChild(backdrop);
        document.body.appendChild(container);
        // Cleanup removes the popup and other elements associated with it.
        const cleanup = () => {
            backdrop.remove();
            container.remove();
            document.body.style.overflow = prevOverflow || "";
            onClose?.();
        };
        backdrop.addEventListener("click", cleanup);
        // Close button
        const closeBtn = container.querySelector(".connective-close-btn");
        if (closeBtn) {
            closeBtn.addEventListener("click", cleanup);
        }
        // Tabs
        const overviewTab = container.querySelector(".connective-tab[data-tab=overview]");
        const configTab = container.querySelector(".connective-tab[data-tab=configuration]");
        const contentEl = container.querySelector(".connective-content-area");
        const setActiveTab = async (tab) => {
            const tabs = container.querySelectorAll(".connective-tab");
            tabs.forEach((t) => t.classList.remove("active"));
            container
                .querySelector(`.connective-tab[data-tab=${tab}]`)
                ?.classList.add("active");
            if (!contentEl)
                return;
            if (tab === "overview") {
                contentEl.innerHTML = await getOverviewHTML(providerData);
            }
            else {
                contentEl.innerHTML = `<p>Configurations are still under development.</p>`;
            }
        };
        overviewTab?.addEventListener("click", () => setActiveTab("overview"));
        configTab?.addEventListener("click", () => setActiveTab("configuration"));
        const connectBtn = container.querySelector(".connective-connect-btn");
        // On-Click handler for "Connect" button
        connectBtn?.addEventListener("click", async () => {
            connectBtn.disabled = true;
            connectBtn.innerText = "Connecting...";
            try {
                const popup = this.openAuthPage(this.projectId, provider, this.userId, this.projectSecret);
                let oauthCompleted = false;
                const onMessage = (event) => {
                    if (event.data?.success) {
                        oauthCompleted = true;
                        setActiveTab("configuration");
                        connectBtn.disabled = false;
                        connectBtn.innerText = "Connected";
                        window.removeEventListener("message", onMessage);
                    }
                };
                window.addEventListener("message", onMessage);
                // Detect manual popup close
                const interval = setInterval(() => {
                    if (popup && popup.closed) {
                        clearInterval(interval);
                        window.removeEventListener("message", onMessage);
                        if (oauthCompleted) {
                            return;
                        }
                        onClose?.();
                        finalCleanup();
                    }
                }, 400);
            }
            catch (err) {
                connectBtn.disabled = false;
                connectBtn.innerText = "Connect";
            }
        });
        // keyboard accessibility: Esc to close
        const escHandler = (e) => {
            if (e.key === "Escape")
                cleanup();
        };
        window.addEventListener("keydown", escHandler);
        // ensure cleanup also removes the key listener
        const originalCleanup = cleanup;
        const finalCleanup = () => {
            window.removeEventListener("keydown", escHandler);
            originalCleanup();
        };
        // override listeners created above to call finalCleanup
        backdrop.onclick = finalCleanup;
        closeBtn && (closeBtn.onclick = finalCleanup);
    }
    close() {
        const el = document.getElementById(this.containerId);
        const backdrop = document.querySelector(".connective-backdrop");
        if (el)
            el.remove();
        if (backdrop)
            backdrop.remove();
        document.body.style.overflow = "";
    }
    openAuthPage(projectID, provider, userID, projectSecret) {
        const url = `${BACKEND_BASE_URL}/oauth/connect?provider=${provider}&projectID=${projectID}&userID=${userID}&projectSecret=${projectSecret}`;
        const popup = window.open(url, "authPopup", "width=500, height=600, left=100, top=100, resizeable=no, scrollbars=yes");
        if (!popup) {
            alert("Please enable popups");
        }
        return popup;
    }
}
//# sourceMappingURL=popup.js.map