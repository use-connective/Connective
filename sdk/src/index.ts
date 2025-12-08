import { type CreatePopupDeps, createPopup } from "./popup/popup.js";
import type { ConnectiveConfig } from "./types.js";

export function createConnective(config: ConnectiveConfig) {
    const deps: CreatePopupDeps = {
        projectId: config.projectId,
        userId: config.userId,
        projectSecret: config.projectSecret,
        baseURL: config.baseURL,
    };

    const popup = createPopup(deps);

    return {
        open: popup.open,
        close: popup.close,
    } as const;
}

export type { ConnectiveConfig };
