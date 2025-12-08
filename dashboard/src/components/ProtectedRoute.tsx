import type { JSX } from "react/jsx-runtime";
import { useNavigate } from "react-router";
import { useEffect, useState } from "react";
import { CheckLoggedInUser } from "@/pages/helpers/user";

const ProtectedRoute = ({ children }: { children: JSX.Element }) => {
    const navigate = useNavigate();
    const [checking, setChecking] = useState(true);

    useEffect(() => {
        const verify = async () => {
            await CheckLoggedInUser(navigate, location.pathname);
            setChecking(false);
        };

        verify();
    }, []);

    if (checking) return null;

    return children;
};

export default ProtectedRoute;
