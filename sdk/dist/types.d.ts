export type OpenArgs = {
    provider: string;
    onSuccess?: (payload: any) => void;
    onClose?: () => void;
};
export type ProviderResponse = {
    message: string;
    data: Provider | null;
    error: string | null;
};
export type Provider = {
    id: number;
    name: string;
    display_name: string;
    auth_type: string;
    image_url: string;
    category: string;
    description: string;
    primary_color: string;
    auth_url: string;
    token_url: string;
    redirect_url: string;
    default_scopes: string[];
    created_at: string;
    updated_at: string;
};
export type ConnectiveConfig = {
    projectId: string;
    userId: string;
    baseURL: string;
    projectSecret: string;
};
//# sourceMappingURL=types.d.ts.map