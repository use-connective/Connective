const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const INTEGRATION = `${BASE_URL}/api/v1/integration`

export const Endpoints = {
    BaseURL: BASE_URL,
    Auth: {
        SignUp: `${BASE_URL}/api/v1/user/create`,
        Login: `${BASE_URL}/api/v1/user/login`,
    },
    Project: {
        Create: `${BASE_URL}/api/v1/project/create`,
        GetAllProjects: `${BASE_URL}/api/v1/project/get-all`,
    },
    CompleteOnboarding: `${BASE_URL}/api/v1/user/complete-onboarding`,
    Integration: {
        GetAllProviders: `${INTEGRATION}/providers`,
        SaveCreds: `${INTEGRATION}/creds/save`,
        GetCreds: `${INTEGRATION}/provider/creds`,
        GetAllCategories: `${INTEGRATION}/provider/categories`,
        GetConnectedAccounts: `${INTEGRATION}/connected-accounts`
    },
};
