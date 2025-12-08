import type { Provider } from "@/api/dto/provider";
import type { APIResponse } from "./dto/api_response";
import { Endpoints } from "@/constants";
import axios from "axios";

export const getProviderById = async (id: number): Promise<Provider> => {
    const resp = await axios.get<APIResponse<Provider>>(
        `${Endpoints.BaseURL}/api/v1/integration/provider`,
        {
            params: { id },
            withCredentials: true,
        }
    );

    return resp.data.data;
};
