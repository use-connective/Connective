import { Button } from "@/components/ui/button";
import { FcGoogle } from "react-icons/fc";
import { FaGithub } from "react-icons/fa";
import { Input } from "@/components/ui/input";
import { Checkbox } from "@/components/ui/checkbox";
import { useState } from "react";
import { useNavigate } from "react-router";
import { handleSignUp } from "./helpers/auth";
import OnboardLayout from "@/components/auth/AuthLayout";

const SignUpPage = () => {
    const navigate = useNavigate();

    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    return (
        <OnboardLayout
            pageType="signup"
            children={
                <div>
                    {/* Social Buttons */}
                    <div className="lg:flex lg:gap-4 mt-8 justify-between">
                        <Button
                            variant="outline"
                            className="h-12 rounded-xl text-sm"
                        >
                            <FcGoogle className="size-5" />
                            <span className="ml-2 px-10 lg:px-4">
                                Sign Up with Google
                            </span>
                        </Button>

                        <div className="mt-4 lg:hidden"></div>

                        <Button
                            variant="outline"
                            className="h-12 rounded-xl text-sm"
                        >
                            <FaGithub className="size-5" />
                            <span className="ml-2 px-10 lg:px-4">
                                Sign Up with GitHub
                            </span>
                        </Button>
                    </div>

                    {/* Divider */}
                    <div className="flex items-center gap-4 my-8">
                        <div className="flex-1 h-px bg-gray-300"></div>
                        <span className="text-gray-500 text-sm">or</span>
                        <div className="flex-1 h-px bg-gray-300"></div>
                    </div>

                    {/* Name */}
                    <div className="mb-5">
                        <label className="font-medium text-gray-800 text-sm lg:text-base">
                            Name
                        </label>
                        <Input
                            placeholder="eg: Sushant"
                            className="mt-2 h-12 rounded-xl placeholder:text-xs placeholder:lg:text-sm"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                    </div>

                    {/* Email */}
                    <div className="mb-5">
                        <label className="font-medium text-gray-800 text-sm lg:text-base">
                            Email
                        </label>
                        <Input
                            placeholder="eg: sushant@example.com"
                            className="mt-2 h-12 rounded-xl placeholder:text-xs placeholder:lg:text-sm"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                        />
                    </div>

                    {/* Password */}
                    <div className="mb-5">
                        <label className="font-medium text-gray-800 text-sm lg:text-base">
                            Password
                        </label>
                        <Input
                            type="password"
                            placeholder="xxxxxxxx"
                            className="mt-2 h-12 rounded-xl placeholder:text-xs placeholder:lg:text-sm"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>

                    {/* Remember */}
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

                    {/* Submit */}
                    <Button
                        className="w-full h-12 text-sm lg:text-base rounded-xl cursor-pointer bg-slate-700 hover:bg-white hover:text-slate-800 border"
                        onClick={() =>
                            handleSignUp(name, email, password, navigate)
                        }
                    >
                        Sign Up
                    </Button>
                </div>
            }
        />
    );
};

export default SignUpPage;
