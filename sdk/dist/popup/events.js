import { openAuthWindow } from "./auth.js";
import { setupTabs } from "./tabs.js";
export function initializePopupEvents(params) {
    const { deps, provider, providerData, container, backdrop, onSuccess, onClose, } = params;
    const { projectId, userId, projectSecret, baseURL } = deps;
    const { setActive } = setupTabs(container, providerData);
    // Buttons
    const closeBtn = container.querySelector(".connective-close-btn");
    const connectBtn = container.querySelector(".connective-connect-btn");
    // ESC handler
    const escHandler = (e) => {
        if (e.key === "Escape")
            finalizeCleanup();
    };
    // message handler
    let oauthCompleted = false;
    const onMessage = (event) => {
        if (event.data?.success) {
            oauthCompleted = true;
            setActive("configuration");
            if (connectBtn) {
                connectBtn.disabled = false;
                connectBtn.innerText = "Connected";
            }
            onSuccess?.(event.data);
            window.removeEventListener("message", onMessage);
        }
    };
    // openAuth flow
    const onConnectClick = async () => {
        if (!connectBtn)
            return;
        connectBtn.disabled = true;
        connectBtn.innerText = "Connecting...";
        const popup = openAuthWindow({
            baseURL,
            projectId,
            provider,
            userId,
            projectSecret,
        });
        window.addEventListener("message", onMessage);
        // poll for close
        const interval = setInterval(() => {
            if (popup && popup.closed) {
                clearInterval(interval);
                window.removeEventListener("message", onMessage);
                if (!oauthCompleted) {
                    onClose?.();
                    finalizeCleanup();
                }
            }
        }, 400);
    };
    // cleanup that removes DOM listeners
    function finalizeCleanup() {
        window.removeEventListener("keydown", escHandler);
        window.removeEventListener("message", onMessage);
        closeBtn?.removeEventListener("click", finalizeCleanup);
        connectBtn?.removeEventListener("click", onConnectClick);
        params.destroy();
    }
    // Wire listeners
    window.addEventListener("keydown", escHandler);
    backdrop.addEventListener("click", finalizeCleanup);
    closeBtn?.addEventListener("click", finalizeCleanup);
    connectBtn?.addEventListener("click", onConnectClick);
    // Return the final cleanup for callers to run
    return finalizeCleanup;
}
//# sourceMappingURL=events.js.map