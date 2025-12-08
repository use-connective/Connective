import { createConnective } from "connective_sdk";

export default function SdkTestPage() {

    const sdk = createConnective({
        projectId: "",
        userId: "",
        baseURL: "http://localhost:8082",
        projectSecret: "",
    });

    const handleConnect = (provider: string) => {
        sdk.open({
            provider,
        });
    };

    return (
        <div className="min-h-screen bg-gray-50 px-4 py-6">
            <div className="max-w-2xl mx-auto">
                {/* Header */}
                <h1 className="text-3xl font-bold text-center mb-6">
                    Connective SDK Test Page
                </h1>

                <p className="text-gray-600 text-center mb-8">
                    Test the Connect Portal from desktop and mobile.
                </p>

                {/* Provider Buttons */}
                <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <button
                        onClick={() => handleConnect("slack")}
                        className="w-full bg-purple-600 hover:bg-purple-700 text-white py-3 rounded-lg shadow-md transition"
                    >
                        Connect Slack
                    </button>

                    <button
                        onClick={() => handleConnect("jira")}
                        className="w-full bg-blue-600 hover:bg-blue-700 text-white py-3 rounded-lg shadow-md transition"
                    >
                        Connect Jira
                    </button>

                    <button
                        onClick={() => handleConnect("microsoft_teams")}
                        className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-3 rounded-lg shadow-md transition"
                    >
                        Connect Microsoft Teams
                    </button>

                    <button
                        onClick={() => handleConnect("github")}
                        className="w-full bg-gray-800 hover:bg-black text-white py-3 rounded-lg shadow-md transition"
                    >
                        Connect GitHub
                    </button>
                </div>
            </div>
        </div>
    );
}
