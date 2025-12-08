import { Endpoints } from "@/constants";
import axios, { isAxiosError } from "axios";
import type { CreateProjectRequest, CreateProjectResponse, Project } from "./dto/project";
import type { APIResponse } from "./dto/api_response";


export const CreateProject = async (name: string) => {
    const req: CreateProjectRequest = { name };

    try {
        const resp = await axios.post<APIResponse<CreateProjectResponse>>(
            Endpoints.Project.Create,
            req,
            {
                withCredentials: true,
            },
        );
        if (resp.data.error) {
            throw new Error(resp.data.error);
        }

        return resp.data;
    } catch (error) {
        if (isAxiosError(error)) {
            const backendMessage =
                error.response?.data?.error ||
                error.response?.data?.message ||
                error.message;

            throw new Error(backendMessage);
        }

        throw new Error("Unexpected error");
    }
};

export const getAllProjects = async (): Promise<Project[]> => {
    try {
        const resp = await axios.get<APIResponse<Project[]>>(Endpoints.Project.GetAllProjects, {
            withCredentials: true,
        });

        if (resp.data.error) {
            throw new Error(resp.data.error);
        }

        return resp.data.data;
    } catch (err) {
        if (isAxiosError(err)) {
            const msg = err.response?.data?.error || err.message;
            throw new Error(msg);
        }

        throw new Error("Unexpected Error");
    }
};