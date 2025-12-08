import type { ConnectiveConfig } from "./types.js";
export declare function createConnective(config: ConnectiveConfig): {
    readonly open: (args: import("./types.js").OpenArgs) => Promise<(() => void) | undefined>;
    readonly close: () => void;
};
export type { ConnectiveConfig };
//# sourceMappingURL=index.d.ts.map