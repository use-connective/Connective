import { Endpoints } from "@/constants";
import axios from "axios";
import type { LoginRequest, LoginResponse, SignUpRequest, SignUpResponse } from "./dto/auth";
import type { APIResponse } from "./dto/api_response";


export const SignUp = async (name: string, email: string, password: string) => {
    const req: SignUpRequest = { name, email, password };

    try {
        const resp = await axios.post<APIResponse<SignUpResponse>>(Endpoints.Auth.SignUp, req, {
            withCredentials: true,
        });
        if (resp.data.error) {
            throw new Error(resp.data.error);
        }

        return resp.data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            const backendMessage =
                error.response?.data?.error ||
                error.response?.data?.message ||
                error.message;

            throw new Error(backendMessage);
        }

        throw new Error("Unexpected error");
    }
};

export const Login = async (email: string, password: string) => {
    const req: LoginRequest = { email, password };

    try {
        const resp = await axios.post<APIResponse<LoginResponse>>(Endpoints.Auth.Login, req, {
            withCredentials: true,
        });
        if (resp.data.error) {
            throw new Error(resp.data.error);
        }

        return resp.data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            const backendMessage =
                error.response?.data?.error ||
                error.response?.data?.message ||
                error.message;

            throw new Error(backendMessage);
        }

        throw new Error("Unexpected error");
    }
};
