import { getOverviewHTML } from "./html.js";
import type { ProviderResponse } from "../types.js";

export function setupTabs(
    container: HTMLElement,
    providerData: ProviderResponse
) {
    const contentEl = container.querySelector<HTMLElement>(
        ".connective-content-area"
    );
    const tabs = Array.from(
        container.querySelectorAll<HTMLElement>(".connective-tab")
    );

    async function setActive(tab: "overview" | "configuration") {
        tabs.forEach((t) => t.classList.remove("active"));
        container
            .querySelector<HTMLElement>(`.connective-tab[data-tab="${tab}"]`)
            ?.classList.add("active");

        if (!contentEl) return;
        if (tab === "overview") {
            contentEl.innerHTML = await getOverviewHTML(providerData);
        } else {
            contentEl.innerHTML = `<p>Configurations are still under development.</p>`;
        }
    }

    container
        .querySelector<HTMLElement>('.connective-tab[data-tab="overview"]')
        ?.addEventListener("click", () => setActive("overview"));
    container
        .querySelector<HTMLElement>('.connective-tab[data-tab="configuration"]')
        ?.addEventListener("click", () => setActive("configuration"));

    return { setActive };
}
