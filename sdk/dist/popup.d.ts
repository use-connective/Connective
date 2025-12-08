import type { OpenArgs } from "./types.js";
export declare class IntegrationPopUpCard {
    private projectId;
    private userId;
    private projectSecret;
    private containerId;
    private styleId;
    constructor(projectId: string, userId: string, projectSecret: string);
    open({ provider, onSuccess, onClose }: OpenArgs): Promise<void>;
    close(): void;
    private openAuthPage;
}
//# sourceMappingURL=popup.d.ts.map