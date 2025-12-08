import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";

import { useNavigate } from "react-router";
import { handleCreateProject } from "./helpers/project";
import OnboardLayout from "@/components/auth/AuthLayout";

const CreateProjectPage = () => {
    const navigate = useNavigate();

    const [project, setProject] = useState("");

    return (
        <OnboardLayout
            pageType="create_project"
            children={
                <div>
                    {/* Email */}
                    <div className="mb-5 mt-8">
                        <label className="font-medium text-gray-800 text-sm lg:text-base">
                            Project Name
                        </label>
                        <Input
                            value={project}
                            onChange={(e) => setProject(e.target.value)}
                            placeholder="eg: My Awesome Project"
                            className="mt-2 h-12 rounded-xl placeholder:text-xs placeholder:lg:text-sm"
                        />
                    </div>

                    {/* Submit Button */}
                    <Button
                        className="w-full h-12 text-sm lg:text-base rounded-xl bg-slate-700 hover:bg-white hover:text-slate-800 border cursor-pointer"
                        onClick={() => {
                            handleCreateProject(project, navigate);
                        }}
                    >
                        Create Project
                    </Button>
                </div>
            }
        />
    );
};

export default CreateProjectPage;
