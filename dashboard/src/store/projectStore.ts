import { create } from "zustand";
import type { Project } from "@/api/dto/project";
import { getAllProjects } from "@/api/project";
import toast from "react-hot-toast";

interface ProjectState {
    projects: Project[];
    selectedProject: Project | null;

    loadProjects: () => Promise<void>;
    setSelectedProject: (project: Project) => void;
}

export const useProjectStore = create<ProjectState>((set) => ({
    projects: [],
    selectedProject: null,

    loadProjects: async () => {
        try {
            const list = await getAllProjects();
            set({
                projects: list,
                selectedProject: list[0] ?? null,
            });
        } catch (err) {
            toast.error("Unable to load project. Please contant support.");
        }
    },

    setSelectedProject: (project: Project) => set({ selectedProject: project }),
}));
