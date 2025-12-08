import { getCSS } from "./css.js";
import { getHTML } from "./html.js";
export async function renderPopupDOM({ providerData, onClose, }) {
    const containerId = "connective-dom-popup-root";
    const styleId = "connective-dom-popup-styles";
    const prevOverflow = document.body.style.overflow;
    document.body.style.overflow = "hidden";
    // inject CSS once
    if (!document.getElementById(styleId)) {
        const style = document.createElement("style");
        style.id = styleId;
        style.innerHTML = getCSS(providerData);
        document.head.appendChild(style);
    }
    // backdrop
    const backdrop = document.createElement("div");
    backdrop.className = "connective-backdrop";
    backdrop.setAttribute("data-connective", "backdrop");
    // container
    const container = document.createElement("div");
    container.id = containerId;
    container.className = "connective-popup-container";
    container.setAttribute("role", "dialog");
    container.setAttribute("aria-modal", "true");
    container.innerHTML = await getHTML(providerData);
    document.body.appendChild(backdrop);
    document.body.appendChild(container);
    const cleanup = () => {
        backdrop.remove();
        container.remove();
        document.body.style.overflow = prevOverflow || "";
        onClose?.();
    };
    return { container, backdrop, cleanup };
}
//# sourceMappingURL=dom.js.map