export interface CreateProjectRequest {
    name: string;
}

export interface CreateProjectResponse {
        created_at: string;
        id: string;
        name: string;
        owner: number;
        updated_at: string;
}


export interface Project {
    id: string;
    name: string;
    owner: number;
    updated_at: string;
    created_at: string;
}
