import { FaSquareGithub } from "react-icons/fa6";
import { IoLogoDiscord } from "react-icons/io5";
import { TbDeviceMobileOff } from "react-icons/tb";
import { FaTwitter } from "react-icons/fa";
import Sidebar from "@/components/dashboard/Sidebar";
import Topbar from "@/components/dashboard/Topbar";
import { Outlet } from "react-router";
import { SiGitconnected } from "react-icons/si";

export const Dashboard = () => {
    return (
        <>
            {/* Mobile View */}
            <MobileView />

            {/* Desktop View */}
            <div className="hidden lg:block">
                <div className="flex h-screen w-screen bg-[#F3F4F7]">
                    <Sidebar />

                    {/* Right Panel */}
                    <div className="flex-1 flex flex-col">
                        <Topbar />

                        {/* Main Content */}
                        <main className="flex-1 p-6 overflow-auto">
                            <Outlet />
                        </main>
                    </div>
                </div>
            </div>
        </>
    );
};

function MobileView() {
    return (
        <div className="md:hidden w-screen h-screen bg-[#F3F4F7] flex flex-col items-center justify-center px-6 text-center">
            {/* Logo  */}

            <div className="flex items-center px-4 py-[1.1rem]">
                <SiGitconnected className="text-gray-800 text-2xl" />

                <p className="text-xl ml-2 font-medium lg:text-lg lg:font-medium text-gray-800">
                    Connective
                </p>
            </div>

            {/* Card Container */}
            <div className="bg-white rounded-2xl border border-gray-200 shadow-sm p-10 max-w-sm w-full flex flex-col items-center">
                {/* Icon */}
                <div className="text-5xl mb-4">
                    <TbDeviceMobileOff />
                </div>

                {/* Heading */}
                <h2 className="text-xl font-semibold text-gray-800">
                    Mobile Not Supported
                </h2>

                {/* Description */}
                <p className="text-sm text-gray-600 mt-3 leading-relaxed">
                    The Connective dashboard is currently optimized for desktop
                    screens. Mobile support is coming soon.
                </p>

                {/* Subtext */}
                <p className="text-sm text-gray-500 mt-4">
                    Want to get notified when mobile view launches?
                </p>

                {/* Notify me button */}
                <button className="mt-6 w-full py-3 rounded-md bg-[#4268FB] text-white text-sm font-medium shadow-sm active:scale-[0.98] transition">
                    Notify Me
                </button>
            </div>

            {/* Footer */}
            <div className="mt-8 flex flex-col items-center gap-3">
                {/* Social Icons */}
                <div className="flex items-center gap-6 text-gray-500">
                    <a
                        href="https://github.com/yourrepo"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <i className="text-2xl">
                            <FaSquareGithub />
                        </i>
                    </a>

                    <a
                        href="https://discord.gg/yourdiscord"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <i className="text-2xl">
                            <IoLogoDiscord />
                        </i>
                    </a>

                    <a
                        href="https://twitter.com/yourhandle"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <i className="text-2xl">
                            <FaTwitter />
                        </i>
                    </a>
                </div>
            </div>
        </div>
    );
}
