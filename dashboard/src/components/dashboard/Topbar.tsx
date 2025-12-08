import {
    DropdownMenu,
    DropdownMenuTrigger,
    DropdownMenuContent,
    DropdownMenuItem,
} from "./../ui/dropdown-menu";

import { NotificationDropdown } from "@/components/dashboard/NotificationPanel";

import { ChevronDown, Gift } from "lucide-react";
import { FaSquareGithub } from "react-icons/fa6";
import { IoLogoDiscord } from "react-icons/io5";
import { useEffect } from "react";
import { useProjectStore } from "@/store/projectStore";

export default function Topbar() {
    const { projects, selectedProject, loadProjects, setSelectedProject } =
        useProjectStore();

    useEffect(() => {
        loadProjects();
    }, []);

    return (
        // <header className="h-16 bg-[#180632] text-white flex items-center justify-between px-6 border-b border-gray-200">
        <header className="h-16 bg-slate-800 text-white flex items-center justify-between px-6 shadow-sm">

        {/* Project Dropdown */}
            <div className="flex items-center gap-2 text-white">
                <p className="font-semibold text-sm ml-1">Project:</p>

                {/* Using shadcn dropdown */}
                <DropdownMenu>
                    <DropdownMenuTrigger className="flex items-center gap-2 cursor-pointer text-sm">
                        <span>{selectedProject?.name ?? "Select Project"}</span>
                        <ChevronDown className="w-4 h-4" />
                    </DropdownMenuTrigger>

                    <DropdownMenuContent>
                        {projects.map((project) => (
                            <DropdownMenuItem
                                key={project.id}
                                onClick={() => setSelectedProject(project)}
                            >
                                {project.name}
                            </DropdownMenuItem>
                        ))}
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>

            {/* Icons + Profile */}
            <div className="flex items-center gap-6">
                <FaSquareGithub className="w-5 h-5 text-white cursor-pointer" />
                <IoLogoDiscord className="w-5 h-5 text-white cursor-pointer" />

                <Gift className="w-5 h-5 text-white cursor-pointer" />
                <NotificationDropdown />

                {/* Profile */}
                <div className="flex items-center gap-3 border-l pl-6">
                    <div className="w-10 h-10 bg-gray-200 rounded-full" />
                    <div className="leading-tight">
                        <p className="font-medium text-sm">Sushant Dhiman</p>
                        <p className="text-xs text-gray-500">
                            sushant.dhiman9812@gmail.com
                        </p>
                    </div>
                </div>
            </div>
        </header>
    );
}
