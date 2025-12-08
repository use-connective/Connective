import { useLocation, useNavigate } from "react-router";
import { Button } from "../ui/button";
import {
    MdOutlineIntegrationInstructions,
    MdOutlineSupport,
    MdOutlineHistory,
} from "react-icons/md";
import { SiGitconnected } from "react-icons/si";
import { AiTwotoneThunderbolt } from "react-icons/ai";
import { LuWorkflow } from "react-icons/lu";
import { SlPeople } from "react-icons/sl";
import { IoDocumentsOutline, IoSettingsOutline } from "react-icons/io5";

export default function Sidebar() {
    const location = useLocation();
    const navigate = useNavigate();
    const path = location.pathname;

    const isExact = (p: string) => path === p;

    const isInSection = (p: string) => path.startsWith(p);

    const activeClass =
        "text-blue-600 font-medium bg-blue-50 rounded-md px-2 py-1";

    const inactiveClass =
        "text-gray-600 hover:text-gray-900 hover:bg-gray-100 rounded-md px-2 py-1";

    const linkClass = (exactPath: string, sectionPath?: string) => {
        if (sectionPath) {
            return isInSection(sectionPath) ? activeClass : inactiveClass;
        }
        return isExact(exactPath) ? activeClass : inactiveClass;
    };

    return (
        <div className="w-64 bg-white text-black flex flex-col">
            {/* Logo */}
            <div
                className="flex items-center bg-slate-800 px-4 py-[1.13rem] cursor-pointer"
                onClick={() => {
                    navigate("/dashboard/providers");
                }}
            >
                <SiGitconnected className="text-white text-2xl" />

                <p className="text-md ml-2 font-bold lg:text-lg lg:font-medium text-white">
                    Connective
                </p>
            </div>

            {/* SECTIONS */}
            <div className="flex flex-col flex-1 border-r border-gray-200">
                {/* INTEGRATE */}
                <div className="px-4 py-8 border-b border-gray-200">
                    <p className="text-xs font-semibold text-gray-500 mb-4">
                        INTEGRATE
                    </p>
                    <div className="flex flex-col gap-2">
                        <div
                            onClick={() => navigate("/dashboard/providers")}
                            className={`flex items-center gap-3 text-sm cursor-pointer ${linkClass(
                                "/dashboard/providers",
                                "/dashboard/providers"
                            )}`}
                        >
                            <MdOutlineIntegrationInstructions />
                            Catalog
                        </div>

                        {/*<div*/}
                        {/*    onClick={() =>*/}
                        {/*        navigate("/dashboard/connected-providers")*/}
                        {/*    }*/}
                        {/*    className={`flex items-center gap-3 text-sm cursor-pointer ${linkClass(*/}
                        {/*        "/dashboard/connected-providers",*/}
                        {/*        "/dashboard/connected-providers"*/}
                        {/*    )}`}*/}
                        {/*>*/}
                        {/*    <SiGitconnected />*/}
                        {/*    Connected Integrations*/}
                        {/*</div>*/}

                        <div
                            onClick={() => navigate("/dashboard/events")}
                            className={`flex items-center gap-3 text-sm cursor-pointer ${linkClass(
                                "/dashboard/events",
                                "/dashboard/events"
                            )}`}
                        >
                            <AiTwotoneThunderbolt />
                            App Events
                        </div>

                        <div
                            onClick={() => navigate("/dashboard/workflow")}
                            className={`flex items-center gap-3 text-sm cursor-pointer ${linkClass(
                                "/dashboard/workflow",
                                "/dashboard/catalog"
                            )}`}
                        >
                            <LuWorkflow /> Workflow
                        </div>
                    </div>
                </div>

                {/* ANALYTICS */}
                <div className="px-4 py-8 border-b border-gray-200">
                    <p className="text-xs font-semibold text-gray-500 mb-4">
                        ANALYTICS
                    </p>
                    <div className="flex flex-col gap-2">
                        <div
                            onClick={() => navigate("/dashboard/users")}
                            className={`flex items-center gap-3 text-sm cursor-pointer ${linkClass(
                                "/dashboard/users",
                                "/dashboard/users"
                            )}`}
                        >
                            <SlPeople /> Connected Users
                        </div>

                        <div
                            onClick={() => navigate("/dashboard/tasks")}
                            className={`flex items-center gap-3 text-sm cursor-pointer ${linkClass(
                                "/dashboard/tasks",
                                "/dashboard/tasks"
                            )}`}
                        >
                            <MdOutlineHistory /> Task History
                        </div>
                    </div>
                </div>

                {/* TUTORIALS */}
                <div className="px-4 py-8 border-b border-gray-200">
                    <p className="text-xs font-semibold text-gray-500 mb-4">
                        TUTORIALS
                    </p>
                    <div className="flex flex-col gap-2">
                        <div
                            onClick={() => navigate("/dashboard/docs")}
                            className={`flex items-center gap-3 text-sm cursor-pointer ${inactiveClass}`}
                        >
                            <IoDocumentsOutline /> Documentation
                        </div>
                    </div>
                </div>

                {/* SUPPORT */}
                <div className="px-4 py-8 border-b border-gray-200">
                    <p className="text-xs font-semibold text-gray-500 mb-4">
                        SUPPORT
                    </p>
                    <div className="flex flex-col gap-2">
                        <div
                            onClick={() => navigate("/dashboard/settings")}
                            className={`flex items-center gap-3 text-sm cursor-pointer ${linkClass(
                                "/dashboard/settings",
                                "/dashboard/settings"
                            )}`}
                        >
                            <IoSettingsOutline /> Settings
                        </div>

                        <div
                            onClick={() => navigate("/dashboard/help")}
                            className={`flex items-center gap-3 text-sm cursor-pointer ${linkClass(
                                "/dashboard/help",
                                "/dashboard/help"
                            )}`}
                        >
                            <MdOutlineSupport /> Support & Help
                        </div>
                    </div>
                </div>
            </div>

            {/* Upgrade */}
            {/*<div className="p-6">*/}
            {/*    <Button className="w-full cursor-pointer bg-slate-700 hover:bg-white hover:text-slate-800 border">*/}
            {/*        Upgrade Plan*/}
            {/*    </Button>*/}
            {/*</div>*/}

            {/* Footer */}
            <p className="text-xs text-center text-gray-500 mt-4 mb-4">
                Connective. Built with ❤️ & Open Source
            </p>
        </div>
    );
}
