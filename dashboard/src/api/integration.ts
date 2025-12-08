import { Endpoints } from "@/constants";
import axios, { isAxiosError } from "axios";
import type { APIResponse } from "./dto/api_response";
import type {
    ConnectedAccount,
    GetCredsRequest,
    IntegrationCreds,
    ProviderDisplayable,
    SaveCredsRequest,
    SaveCredsResponse,
} from "./dto/integration";

export const getAllProviders = async (search: string, category: string): Promise<ProviderDisplayable[]> => {
    try {
        let endpoint = Endpoints.Integration.GetAllProviders + '?'

        if (search !== undefined && search !== '') {
            endpoint += `search=${search}&`
        }

        if (category !== undefined && category !== '') {
            endpoint += `category=${category}`
        }

        const resp = await axios.get<APIResponse<ProviderDisplayable[]>>(
            endpoint,
            { withCredentials: true }
        );

        if (resp.data.error) {
            throw new Error(resp.data.error);
        }

        const data = resp.data.data;
        return Array.isArray(data) ? data : [data];
    } catch (err) {
        if (isAxiosError(err)) {
            const msg = err.response?.data?.error || err.message;
            throw new Error(msg);
        }
        throw new Error("Unexpected Error");
    }
};

export const saveIntegrationCreds = async (req: SaveCredsRequest) => {
    try {
        const resp = await axios.post<APIResponse<SaveCredsResponse>>(
            Endpoints.Integration.SaveCreds,
            req,
            {
                withCredentials: true,
            }
        );

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

export const getIntegrationCreds = async (req: GetCredsRequest) => {
    try {
        const resp = await axios.get<APIResponse<IntegrationCreds>>(
            `${Endpoints.Integration.GetCreds}?providerID=${req.providerID}&projectID=${req.projectID}`,
            {
                withCredentials: true,
            }
        );

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

export const getAllCategories = async () => {
    try {
        const resp = await axios.get<APIResponse<string[]>>(
            Endpoints.Integration.GetAllCategories,
            {
                withCredentials: true,
            }
        );

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

export const getConnectedAccounts = async (projectId: string, userId: string) => {
    try {
        let endpoint = Endpoints.Integration.GetConnectedAccounts + `?projectId=${projectId}&userId=${userId}`

        const resp = await axios.get<APIResponse<ConnectedAccount[]>>(
            endpoint,
            {
                withCredentials: true,
            }
        );

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