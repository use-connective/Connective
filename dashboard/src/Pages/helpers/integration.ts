import { toast } from "react-hot-toast";
import type { NavigateFunction } from "react-router";
import type { SaveCredsRequest } from "@/api/dto/integration";
import { saveIntegrationCreds } from "@/api/integration";

export const handleSaveIntegrationCreds = async (
    creds: SaveCredsRequest,
    navigate: NavigateFunction,
    redirectPath: string = "/integrations" // default redirect
) => {
    if (!creds.client_id || !creds.client_secret || !creds.project_id) {
        toast.error("Client ID, secret, and project must be provided.");
        return;
    }

    return toast
        .promise(saveIntegrationCreds(creds), {
            loading: "Saving credentials...",
            success: "Credentials saved successfully!",
            error: (err) => err.message || "Failed to save credentials",
        })
        // .then(() => navigate(redirectPath))
        .catch(() => {});
};
