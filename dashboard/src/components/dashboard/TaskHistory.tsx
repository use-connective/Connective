import { Info } from "lucide-react";
import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";

export default function TaskHistory() {
    return (
        <div className="bg-white border rounded-2xl p-8 h-full flex flex-col items-center justify-center text-center">
            <div className="bg-blue-50 text-slate-700 p-4 rounded-full mb-4">
                <Info className="w-8 h-8" />
            </div>

            <h2 className="text-2xl font-semibold text-gray-900 mb-2">
                Task History Coming Soon
            </h2>

            <p className="text-gray-600 max-w-md mb-6">
                We're actively building this feature. Join our Discord community to get updates,
                vote on features, and be part of the roadmap.
            </p>

            <a
                // TODO - Replace with actual Discord Link
                href="https://discord.gg/"
                target="_blank"
                rel="noopener noreferrer"
            >
                <Button className="flex items-center gap-2 px-5 py-2 bg-slate-700">
                    Join Discord
                    <ArrowRight className="w-4 h-4" />
                </Button>
            </a>
        </div>
    );
}
