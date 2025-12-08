import { fetchProvider } from "../api.js";
import { renderPopupDOM } from "./dom.js";
import { initializePopupEvents } from "./events.js";
import type { OpenArgs } from "../types.js";

export type CreatePopupDeps = {
    baseURL: string;
    projectId: string;
    userId: string;
    projectSecret: string;
};

export function createPopup(deps: CreatePopupDeps) {
    const containerId = "connective-dom-popup-root";

    async function open(args: OpenArgs) {
        if (document.getElementById(containerId)) return;

        const providerData = await fetchProvider(deps.baseURL, args.provider);

        const { container, backdrop, cleanup } = await renderPopupDOM({
            providerData,
            onClose: args.onClose,
        });

        const finalCleanup = initializePopupEvents({
            deps,
            provider: args.provider,
            providerData,
            container,
            backdrop,
            onSuccess: args.onSuccess,
            onClose: args.onClose,
            destroy: () => {
                cleanup()
            }
        });

        return () => {
            finalCleanup();
            cleanup();
        };
    }

    function close() {
        const el = document.getElementById(containerId);
        const backdrop = document.querySelector(".connective-backdrop");
        el?.remove();
        backdrop?.remove();
        document.body.style.overflow = "";
    }

    return { open, close } as const;
}
