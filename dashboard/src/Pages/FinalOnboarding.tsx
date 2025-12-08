import { Button } from "@/components/ui/button";
import { FaDiscord } from "react-icons/fa";
import { toast } from "react-hot-toast";
import { useNavigate } from "react-router";
import { CompleteUserOnboarding } from "@/api/final_onboarding";
import OnboardLayout from "@/components/auth/AuthLayout";

const FinalOnboardingPage = () => {
    const navigate = useNavigate();

    async function handleSubmit() {
        toast
            .promise(CompleteUserOnboarding(), {
                success: "Onboarding Completed",
                error: (err) => err.message,
                loading: "Completing Onboarding",
            })
            .then(() => navigate("/dashboard/providers"))
            .catch(() => {});
    }

    return (
        <OnboardLayout
            pageType="final_onboarding"
            children={
                <div>
                    <p className="text-justify text-gray-600 mt-3 text-sm lg:text-md">
                        I’m really excited to see you complete all the
                        onboarding steps — that means you’re all set to start
                        integrating platforms. <br /> <br /> Before you dive in,
                        I’d love to invite you to join our Discord community.
                        It’s the place where I connect directly with users,
                        understand their real needs, and shape the product
                        around your workflow. Your feedback will have a direct
                        impact on what we build next, and I want you to be part
                        of that journey from day one.
                    </p>

                    {/* Social Buttons */}
                    <div className="lg:flex lg:gap-4 mt-8 justify-between">
                        <Button
                            variant="outline"
                            className="h-12 rounded-xl text-sm cursor-pointer"
                        >
                            <FaDiscord className="size-5" />
                            <span className="ml-2 px-16 lg:px-8">
                                Join Discord
                            </span>
                        </Button>

                        <div className="mt-4 lg:hidden"></div>

                        <Button
                            variant="default"
                            className="h-12 rounded-xl text-sm bg-slate-700 hover:bg-white hover:text-slate-800 border cursor-pointer"
                            onClick={handleSubmit}
                        >
                            <span className="ml-2 px-12.5 lg:px-4 cursor-pointer">
                                Explore Integrations
                            </span>
                        </Button>
                    </div>
                </div>
            }
        />
    );
};

export default FinalOnboardingPage;
