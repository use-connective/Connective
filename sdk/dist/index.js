import { createPopup } from "./popup/popup.js";
export function createConnective(config) {
    const deps = {
        projectId: config.projectId,
        userId: config.userId,
        projectSecret: config.projectSecret,
        baseURL: config.baseURL,
    };
    const popup = createPopup(deps);
    return {
        open: popup.open,
        close: popup.close,
    };
}
//# sourceMappingURL=index.js.map