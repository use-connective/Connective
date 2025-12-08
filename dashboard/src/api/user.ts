import axios, { isAxiosError } from "axios";

const USER_BASE_URL = "http://localhost:8082/api/v1/user/";

interface User {
  message: string,
  data: {
    id: number,
    name: string,
    email: string,
    created_at: string,
    updated_at: string,
    is_onboarding_completed: boolean,
    state: string
  },
  error: string
}

export const GetLoggedInUser = async () => {
    try {
        const resp = await axios.get<User>(USER_BASE_URL,
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
