import { useNavigate, useParams } from "react-router";
import { useEffect, useState } from "react";
import { getProviderById } from "@/api/provider";
import type { Provider } from "@/api/dto/provider";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import toast from "react-hot-toast";
import { useProjectStore } from "@/store/projectStore";
import { handleSaveIntegrationCreds } from "@/pages/helpers/integration";
import { FaRegCopy } from "react-icons/fa";
import { IoCloseCircleOutline } from "react-icons/io5";
import { createConnective } from "connective_sdk";
import { getIntegrationCreds } from "@/api/integration";

export default function ProviderDetails() {
    const { id } = useParams();
    const [provider, setProvider] = useState<Provider | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();
    const project = useProjectStore((state) => state.selectedProject);

    const [clientId, setClientId] = useState("");
    const [clientSecret, setClientSecret] = useState("");
    const [scopes, setScopes] = useState<string>("");

    const sdk = createConnective({
        projectId: project?.id ?? '',
        userId: "1",
        baseURL: "http://localhost:8082",
        projectSecret: 'aaca6a2a-f3ce-4f76-963a-35cc9fe5e7c0',
    });

    const copyToClipboard = async () => {
        await navigator.clipboard.writeText(provider?.redirect_url ?? "");
        toast.success("Copied to clipboard.");
    };

    useEffect(() => {
        if (!id) {
            setError("No provider ID provided");
            setLoading(false);
            return;
        }

        const loadProvider = async () => {
            try {
                setLoading(true);
                setError(null);

                const data = await getProviderById(Number(id));
                setProvider(data);
                setScopes(data.default_scopes.join(","));
            } catch (err: any) {
                setError("Failed to load provider details");
                toast.error(err?.message ?? "Failed to load provider");
            } finally {
                setLoading(false);
            }
        };

        const loadIntegrationCreds = async () => {
            if (!project) return;

            try {
                const data = await getIntegrationCreds({
                    projectID: project.id,
                    providerID: id,
                });

                if (data.data != null) {
                    setClientId(data.data?.client_id ?? "");
                    setClientSecret(data.data?.client_secret ?? "");
                    setScopes(data.data?.scopes.join(",") ?? "");
                }
            } catch (err: any) {
                toast.error(err?.message ?? "Failed to load credentials");
            }
        };

        loadProvider();
        loadIntegrationCreds();
    }, [id, project]);

    if (loading) {
        return (
            <div className="flex items-center justify-center h-full">
                <p className="text-gray-600">Loading...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div className="flex flex-col items-center justify-center h-full gap-4">
                <div className="text-6xl">❌</div>
                <p className="text-xl font-semibold text-gray-800">
                    Error Occurred
                </p>
                <p className="text-gray-600">{error}</p>
            </div>
        );
    }

    if (!provider) {
        return (
            <div className="flex flex-col items-center justify-center h-full gap-4">
                <div className="text-6xl">⚠️</div>
                <p className="text-xl font-semibold text-gray-800">
                    Provider Not Found
                </p>
            </div>
        );
    }

    // ---------- MAIN UI ----------
    return (
        <div className="bg-white border rounded-2xl p-6 h-full">
            <div className="w-xl">
                <div className="flex items-center justify-between">
                    <h2 className="text-xl font-semibold mb-4">
                        Configure Provider: {provider.display_name}
                    </h2>

                    <div className="flex space-x-4">
                        <p
                            className="text-blue-600 font-medium cursor-pointer text-sm"
                            onClick={() => navigate("/dashboard/providers")}
                        >
                            How To?
                        </p>
                        <p
                            className="cursor-pointer text-sm"
                            onClick={() => navigate("/dashboard/providers")}
                        >
                            <IoCloseCircleOutline className="text-2xl" />
                        </p>
                    </div>
                </div>

                <div className="flex flex-col space-y-4">
                    {/* Client ID */}
                    <div className="flex flex-col gap-2">
                        <label className="text-sm">oAuth Client ID</label>
                        <Input
                            className="mt-1 border rounded p-2 w-full"
                            value={clientId}
                            onChange={(e) => setClientId(e.target.value)}
                        />
                    </div>

                    {/* Client Secret */}
                    <div className="flex flex-col gap-2">
                        <label className="text-sm">oAuth Client Secret</label>
                        <Input
                            className="mt-1 border rounded p-2 w-full"
                            value={clientSecret}
                            onChange={(e) => setClientSecret(e.target.value)}
                        />
                    </div>

                    {/* Scopes */}
                    <div className="flex flex-col gap-2">
                        <label className="text-sm">
                            oAuth Scopes (Comma Separated)
                        </label>
                        <Input
                            className="mt-1 border rounded p-2 w-full"
                            value={scopes}
                            onChange={(e) => setScopes(e.target.value)}
                        />
                    </div>

                    {/* Redirect URL */}
                    <div className="flex flex-col gap-2 relative">
                        <label className="text-sm">Redirect URL</label>

                        <div className="relative w-full">
                            <Input
                                disabled
                                className="mt-1 border rounded p-2 w-full pr-10"
                                value={provider.redirect_url}
                            />

                            <button
                                type="button"
                                onClick={copyToClipboard}
                                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-600 hover:text-black"
                            >
                                <FaRegCopy className="cursor-pointer" />
                            </button>
                        </div>
                    </div>

                    {/* Buttons */}
                    <div className="flex space-x-4 mt-4">
                        <Button
                            variant={"outline"}
                            className="cursor-pointer"
                            onClick={() => {
                                sdk.open({
                                    provider: provider.name,
                                });
                            }}
                        >
                            Test Integration
                        </Button>

                        <Button
                            className="cursor-pointer bg-slate-700"
                            onClick={() => {
                                if (!project) {
                                    toast.error(
                                        "Please select a project first."
                                    );
                                    return;
                                }

                                const creds = {
                                    client_id: clientId,
                                    client_secret: clientSecret,
                                    provider_id: provider.id,
                                    project_id: project.id.toString(),
                                    scopes: scopes
                                        .split(",")
                                        .map((s) => s.trim())
                                        .filter(Boolean),
                                };

                                handleSaveIntegrationCreds(
                                    creds,
                                    navigate,
                                    "/dashboard/providers"
                                );
                            }}
                        >
                            Save Credentials
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    );
}
