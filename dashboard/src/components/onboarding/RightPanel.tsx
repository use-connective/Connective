import DashboardV2 from "../../assets/DashboardV2.png";
import DashboardV3 from "../../assets/DashboardV3.png";
import {useEffect, useState} from "react";


export const slides = [
    {
        title: "The Complete Integration Platform\nFor Your SaaS",
        description:
            "Connective gives SaaS teams everything they need to ship integrations fast—OAuth flows, token refresh, syncing, events, and webhooks—without building infrastructure from scratch. Focus your time on core product development while Connective automates the heavy lifting behind every connection.",
        image: DashboardV2,
    },
    {
        title: "Instant Plug-and-Play Integrations\nBuilt for Modern SaaS",
        description:
            "Enable your users to instantly connect with the tools they already rely on—like Slack, Google Calendar, HubSpot, and more. Connective provides a unified integration layer that handles authentication, data mapping, background syncing, and error handling out of the box.",
        image: DashboardV3,
    },
    {
        title: "Accelerate Your Roadmap\nWith Zero Maintenance",
        description:
            "Stop maintaining dozens of brittle, one-off integrations. Connective provides a scalable architecture, shared OAuth logic, version-safe connectors, and a single API interface. Spend engineering effort on features that differentiate your SaaS—not integration plumbing.",
        image: DashboardV2,
    },
];

function OnboardRightPanel() {
    const [current, setCurrent] = useState(0);

    useEffect(() => {
        const interval = setInterval(() => {
            setCurrent((prev) => (prev + 1) % slides.length);
        }, 5000);

        return () => clearInterval(interval);
    }, []);

    return (
        <div className="w-full overflow-hidden">
            <div
                className="flex transition-transform duration-700"
                style={{ transform: `translateX(-${current * 100}%)` }}
            >
                {slides.map((slide, i) => (
                    <div
                        key={i}
                        className="min-w-full flex flex-col items-center text-center"
                    >
                        <p className="text-white font-medium text-3xl mt-18 whitespace-pre-line">
                            {slide.title}
                        </p>

                        <div className="p-24">
                            <img src={slide.image} className="border-2 rounded-md" />
                        </div>

                        <p className="text-white text-md px-24">{slide.description}</p>
                    </div>
                ))}
            </div>

            {/* Dots */}
            <div className="flex justify-center gap-2 mt-6">
                {slides.map((_, idx) => (
                    <button
                        key={idx}
                        onClick={() => setCurrent(idx)}
                        className={`w-2 h-2 rounded-full transition-all ${
                            current === idx ? "bg-white" : "bg-gray-500"
                        } cursor-pointer`}
                    />
                ))}
            </div>
        </div>
    );
}


export default OnboardRightPanel
