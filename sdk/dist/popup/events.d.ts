import type { CreatePopupDeps } from "./popup.js";
import type { ProviderResponse } from "../types.js";
export declare function initializePopupEvents(params: {
    deps: CreatePopupDeps;
    provider: string;
    providerData: ProviderResponse;
    container: HTMLElement;
    backdrop: HTMLElement;
    onSuccess?: ((payload: any) => void) | undefined;
    onClose?: (() => void) | undefined;
    destroy: () => void;
}): () => void;
//# sourceMappingURL=events.d.ts.map