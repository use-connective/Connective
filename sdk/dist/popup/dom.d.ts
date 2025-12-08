import type { ProviderResponse } from "../types.js";
export declare function renderPopupDOM({ providerData, onClose, }: {
    providerData: ProviderResponse;
    onClose?: (() => void) | undefined;
}): Promise<{
    readonly container: HTMLDivElement;
    readonly backdrop: HTMLDivElement;
    readonly cleanup: () => void;
}>;
//# sourceMappingURL=dom.d.ts.map