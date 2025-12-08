import { getOverviewHTML } from "./html.js";
export function setupTabs(container, providerData) {
    const contentEl = container.querySelector(".connective-content-area");
    const tabs = Array.from(container.querySelectorAll(".connective-tab"));
    async function setActive(tab) {
        tabs.forEach((t) => t.classList.remove("active"));
        container
            .querySelector(`.connective-tab[data-tab="${tab}"]`)
            ?.classList.add("active");
        if (!contentEl)
            return;
        if (tab === "overview") {
            contentEl.innerHTML = await getOverviewHTML(providerData);
        }
        else {
            contentEl.innerHTML = `<p>Configurations are still under development.</p>`;
        }
    }
    container
        .querySelector('.connective-tab[data-tab="overview"]')
        ?.addEventListener("click", () => setActive("overview"));
    container
        .querySelector('.connective-tab[data-tab="configuration"]')
        ?.addEventListener("click", () => setActive("configuration"));
    return { setActive };
}
//# sourceMappingURL=tabs.js.map