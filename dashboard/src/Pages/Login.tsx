import { Button } from "@/components/ui/button";
import { FcGoogle } from "react-icons/fc";
import { FaGithub } from "react-icons/fa";
import { Input } from "@/components/ui/input";
import { Checkbox } from "@/components/ui/checkbox";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import { handleLogin } from "./helpers/auth";
import { useLocation } from "react-router";
import toast from "react-hot-toast";
import OnboardLayout from "@/components/auth/AuthLayout";

const LoginPage = () => {
    const navigate = useNavigate();
    const location = useLocation();

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    useEffect(() => {
        if (location.state?.fromProtectedRoute) {
            toast.error("You are not logged in. Please login again.");
        }
    }, []);

    return (
        <OnboardLayout
            pageType="login"
            children={
                <div>
                    {/* Social buttons */}
                    <div className="lg:flex lg:gap-4 mt-8 justify-between">
                        <Button
                            variant="outline"
                            className="h-12 rounded-xl text-sm"
                        >
                            <FcGoogle className="size-5" />
                            <span className="ml-2 px-10 lg:px-4">
                                Login with Google
                            </span>
                        </Button>

                        <div className="mt-4 lg:hidden"></div>

                        <Button
                            variant="outline"
                            className="h-12 rounded-xl text-sm"
                        >
                            <FaGithub className="size-5" />
                            <span className="ml-2 px-10 lg:px-4">
                                Login with GitHub
                            </span>
                        </Button>
                    </div>

                    {/* Divider */}
                    <div className="flex items-center gap-4 my-8">
                        <div className="flex-1 h-px bg-gray-300"></div>
                        <span className="text-gray-500 text-sm">or</span>
                        <div className="flex-1 h-px bg-gray-300"></div>
                    </div>

                    {/* Email */}
                    <div className="mb-5">
                        <label className="font-medium text-gray-800 text-sm lg:text-base">
                            Email
                        </label>
                        <Input
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="eg: sushant.dhiman9812@gmail.com"
                            className="mt-2 h-12 rounded-xl placeholder:text-xs placeholder:lg:text-sm"
                        />
                    </div>

                    {/* Password */}
                    <div className="mb-5">
                        <label className="font-medium text-gray-800 text-sm lg:text-base">
                            Password
                        </label>
                        <Input
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            type="password"
                            placeholder="xxxxxxxxx"
                            className="mt-2 h-12 rounded-xl placeholder:text-xs placeholder:lg:text-sm"
                        />
                    </div>

                    {/* Remember / Forgot */}
                    <div className="flex items-center justify-between mb-8">
                        <label className="flex items-center space-x-2 text-gray-800">
                            <Checkbox id="remember" />
                            <span className="text-xs lg:text-sm">
                                Remember Me
                            </span>
                        </label>

                        <button className="text-xs lg:text-sm text-blue-600 hover:underline">
                            Forgot Password?
                        </button>
                    </div>

                    {/* Submit button */}
                    <Button
                        className="w-full h-12 text-sm lg:text-base rounded-xl bg-slate-700 hover:bg-white hover:text-slate-800 border cursor-pointer"
                        onClick={() => handleLogin(email, password, navigate)}
                    >
                        Login
                    </Button>
                </div>
            }
        />
    );
};

export default LoginPage;
