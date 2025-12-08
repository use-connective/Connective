export interface SignUpRequest {
    email: string;
    name: string;
    password: string;
}

export interface SignUpResponse {
    token: string;
    user: {
        created_at: string;
        email: string;
        id: number;
        name: string;
        password: string;
        updated_at: string;
    };
}

export interface LoginRequest {
    email: string;
    password: string;
}

export interface LoginResponse {
    token: string;
    user: {
        created_at: string;
        email: string;
        id: number;
        name: string;
        password: string;
        updated_at: string;
    };
}
