import { GetLoggedInUser } from "@/api/user";
import type { NavigateFunction } from "react-router";

export const CheckLoggedInUser = async (
    navigate: NavigateFunction,
    currentPath: string
) => {
    try {
        const user = await GetLoggedInUser();

        const state = user.data?.state;

        if (state === "PROJECT_PENDING") {
            navigate("/create-project");
            return;
        }

        if (state === "ONBOARDING_COMPLETION_PENDING") {
            navigate("/complete-onboarding");
            return;
        }

        if (state === "ACTIVE") {
            if (!currentPath.startsWith("/dashboard")) {
                navigate("/dashboard");
            }
        }
    } catch {
        navigate("/login", { state: { fromProtectedRoute: true } });
    }
};
