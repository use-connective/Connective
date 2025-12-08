import type { OpenArgs } from "../types.js";
export type CreatePopupDeps = {
    baseURL: string;
    projectId: string;
    userId: string;
    projectSecret: string;
};
export declare function createPopup(deps: CreatePopupDeps): {
    readonly open: (args: OpenArgs) => Promise<(() => void) | undefined>;
    readonly close: () => void;
};
//# sourceMappingURL=popup.d.ts.map