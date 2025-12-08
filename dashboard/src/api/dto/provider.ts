export interface Provider {
    id: number;
    name: string;
    display_name: string;
    auth_type: string;
    image_url: string;
    category: string;
    auth_url: string;
    token_url: string;
    redirect_url: string;
    default_scopes: string[];
    created_at: string;
    updated_at: string;
}
