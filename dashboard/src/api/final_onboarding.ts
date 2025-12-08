import { Endpoints } from "@/constants";
import axios, { isAxiosError } from "axios";
import type { APIResponse } from "./dto/api_response";

export const CompleteUserOnboarding = async () => {
    try {
        const resp = await axios.post<APIResponse<string>>(Endpoints.CompleteOnboarding, null,
            {
                withCredentials: true,
            },
        );

        if (resp.data.error) {
            return Promise.reject(new Error(resp.data.error));
        }

        return resp.data;
    } catch (err) {
        let message;

        if (isAxiosError(err)) {
            message = err.response?.data?.error || err.response?.data?.message || err.message;
        } else {
            message = "Unexpected Error.";
        }

        return Promise.reject(new Error(message));
    }
};
