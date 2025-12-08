import type React from "react";
import OnboardRightPanel from "../onboarding/RightPanel";
import { SiGitconnected } from "react-icons/si";

type PageType = "signup" | "login" | "create_project" | "final_onboarding";

type AuthLayoutProps = {
    children: React.ReactNode;
    pageType: PageType;
};

const TextData = {
    SignUp: {
        title: "Get Started",
        subtitle: "Create an account and start using Connective.",
    },
    Login: {
        title: "Welcome Back",
        subtitle: "Please enter your information to access your account.",
    },
    CreateProject: {
        title: "Create Project",
        subtitle: "It’s time to create a project. Give it a good name.",
    },
    FinalOnboarding: {
        title: "Welcome To Connective 👋",
        subtitle: "",
    },
};

function getTitle(pageType: PageType) {
    switch (pageType) {
        case "login":
            return TextData.Login.title;
        case "signup":
            return TextData.SignUp.title;
        case "create_project":
            return TextData.CreateProject.title;
        case "final_onboarding":
            return TextData.FinalOnboarding.title;
    }
}

function getSubtitle(pageType: PageType) {
    switch (pageType) {
        case "login":
            return TextData.Login.subtitle;
        case "signup":
            return TextData.SignUp.subtitle;
        case "create_project":
            return TextData.CreateProject.subtitle;
        case "final_onboarding":
            return TextData.FinalOnboarding.subtitle;
    }
}

export default function OnboardLayout(props: AuthLayoutProps) {
    return (
        <div className="w-screen min-h-screen h-screen lg:h-screen bg-[#f5f5f5] flex flex-col lg:flex-row p-4 gap-4">
            {/* LEFT PANEL */}
            <div className="bg-white w-full lg:w-1/2 h-full rounded-2xl border border-gray-300 px-10 flex flex-col">
                {/* Header */}
                <div className="flex items-center justify-between pt-4">
                    {/* <LogoAndTitle Logo={Logo} Width={60} /> */}

                    <div className="flex items-center px-4 py-[1.1rem]">
                        <SiGitconnected className="text-gray-800 text-3xl" />

                        <p className="text-xl ml-2 font-medium lg:text-2xl lg:font-medium text-gray-800">
                            Connective
                        </p>
                    </div>

                    {props.pageType === "login" ||
                    props.pageType === "signup" ? (
                        <div className="flex items-center space-x-1 text-gray-700">
                            <span className="hidden lg:block lg:text-sm">
                                {props.pageType === "signup"
                                    ? "Already have an account?"
                                    : "Don't have an account?"}
                            </span>
                            <a
                                href={
                                    props.pageType == "signup"
                                        ? "/login"
                                        : "sign-up"
                                }
                                className="font-medium hover:underline"
                            >
                                <span className="text-xs font-bold lg:text-sm">
                                    {props.pageType === "signup"
                                        ? "Login"
                                        : "Sign Up"}
                                </span>
                            </a>
                        </div>
                    ) : (
                        <></>
                    )}
                </div>

                <div className="flex flex-col justify-between flex-1">
                    {/* Main form area */}
                    <div className="flex justify-center pt-8 lg:pt-16 h-full">
                        <div className="w-full max-w-md px-2 flex flex-col">
                            {/* Heading */}
                            <p className="text-xl lg:text-3xl font-medium">
                                {getTitle(props.pageType)}
                            </p>
                            <p className="text-gray-600 mt-3 text-sm lg:text-md">
                                {getSubtitle(props.pageType)}
                            </p>

                            {props.children}
                        </div>
                    </div>

                    {/* Footer — fixed to bottom */}
                    <p className="text-xs text-center mt-10 mb-4">
                        <b>Connective</b>. Developed with ❤️ and Open Source.
                    </p>
                </div>
            </div>

            {/* RIGHT PANEL */}
            <div className="hidden lg:block bg-slate-700 w-1/2 h-full rounded-2xl border border-gray-300">
                <OnboardRightPanel />
            </div>
        </div>
    );
}
