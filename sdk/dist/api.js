import axios from "axios";
/**
 * fetchProvider accepts baseURL so there's no global.
 */
export async function fetchProvider(baseURL, provider) {
    try {
        const res = await axios.get(`${baseURL}/api/v1/integration/provider/get-provider`, {
            params: { provider },
            headers: { accept: "application/json" },
            withCredentials: true,
        });
        return {
            message: res.data?.message ?? "Success",
            data: res.data?.data ?? null,
            error: null,
        };
    }
    catch (err) {
        const status = err?.response?.status;
        const message = err?.response?.data?.message;
        const errorText = status
            ? `HTTP Error: ${status}`
            : err?.message ?? "Unknown error";
        return {
            message: message ?? "Failed to fetch provider",
            data: null,
            error: errorText,
        };
    }
}
//# sourceMappingURL=api.js.map