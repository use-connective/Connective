import { toast } from "react-hot-toast";
import { SignUp, Login } from "@/api/auth";
import type { NavigateFunction } from "react-router";

export const handleSignUp = (
    name: string,
    email: string,
    password: string,
    navigate: NavigateFunction
) => {
    if (name === undefined || name === "") {
        toast.error("Name must be provided.");
        return;
    }

    if (email === undefined || email === "") {
        toast.error("Email must be provided.");
        return;
    }

    if (password === undefined || password === "") {
        toast.error("Password must be provided.");
        return;
    }

    toast
        .promise(SignUp(name, email, password), {
            success: "Account Created",
            error: (err) => err.message,
            loading: "Creating Account",
        })
        .then(() => navigate("/create-project"))
        .catch(() => {});
};

export const handleLogin = async (
    email: string,
    password: string,
    navigate: NavigateFunction
) => {
    if (email === undefined || email === "") {
        toast.error("Email must be provided.");
        console.log("s");
        return;
    }

    if (password === undefined || password === "") {
        toast.error("Password must be provided.");
        return;
    }

    toast
        .promise(Login(email, password), {
            success: "Login Successful",
            error: (err) => err.message,
            loading: "Logging In",
        })
        .then(() => navigate("/dashboard/providers"))
        .catch(() => {});
};
