import { toast } from "react-hot-toast";
import { CreateProject } from "@/api/project";
import type { NavigateFunction } from "react-router";

export const handleCreateProject = (
    project: string,
    navigate: NavigateFunction,
) => {
    if (project === undefined || project === "") {
        toast.error("Project name must be provided.");
        return;
    }

    toast
        .promise(CreateProject(project), {
            success: "Project Created",
            error: (err) => err.message,
            loading: "Creating Project",
        })
        .then(() => navigate("/complete-onboarding"))
        .catch(() => {});
};
