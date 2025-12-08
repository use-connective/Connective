export interface ProviderDisplayable {
    id: number;
    name: string;
    display_name: string;
    image_url: string;
}

export interface SaveCredsRequest {
    client_id: string;
    client_secret: string;
    project_id: string;
    provider_id: number;
    scopes: string[];
}

export interface SaveCredsResponse {
    id: string;
    created_at: string;
    updated_at: string;
    client_id: string;
    provider_id: number;
    project_id: string;
}

export type GetCredsRequest = {
    providerID: string;
    projectID: string;
};

export interface IntegrationCreds {
    id: number;
    project_id: string;
    provider_id: number;
    client_id: string;
    client_secret: string;
    scopes: string[];
    is_active: boolean;
    created_at: string;
    updated_at: string;
}

export interface ConnectedAccount {
    user_id: string
    integrations_enabled: string
    displayable_date: string
}