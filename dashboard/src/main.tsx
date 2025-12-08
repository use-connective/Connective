import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter, Route, Routes } from "react-router";
import "@fontsource/inter/400.css";
import "@fontsource/inter/500.css";
import "@fontsource/inter/600.css";
import "@fontsource/inter/700.css";
import SignUpPage from "./pages/SignUp.tsx";
import LoginPage from "./pages/Login.tsx";
import CreateProjectPage from "./pages/CreateProject.tsx";
import { Toaster } from "react-hot-toast";
import FinalOnboardingPage from "./pages/FinalOnboarding.tsx";
import { Dashboard } from "./pages/Dashboard.tsx";
import ProtectedRoute from "./components/ProtectedRoute.tsx";
import ProvidersList from "./components/dashboard/ProvidersList.tsx";
import ProviderDetails from "./components/dashboard/ProviderDetails.tsx";
import SdkTestPage from "./test/ConnectPortalTest.tsx";
import AppEvents from "@/components/dashboard/AppEvents.tsx";
import Workflow from "@/components/dashboard/Workflow.tsx";
import TaskHistory from "@/components/dashboard/TaskHistory.tsx";
import ConnectedAccounts from "@/components/dashboard/ConnectedAccounts.tsx";

createRoot(document.getElementById("root")!).render(
    <BrowserRouter>
        <Toaster position="bottom-right" reverseOrder={false} />

        <StrictMode>
            <Routes>
                <Route index element={<App />} />
                <Route path="/sign-up" element={<SignUpPage />} />
                <Route path="/login" element={<LoginPage />} />

                <Route
                    path="/create-project"
                    element={
                        <ProtectedRoute>
                            <CreateProjectPage />
                        </ProtectedRoute>
                    }
                />

                <Route
                    path="/complete-onboarding"
                    element={
                        <ProtectedRoute>
                            <FinalOnboardingPage />
                        </ProtectedRoute>
                    }
                />

                <Route
                    path="/dashboard"
                    element={
                        <ProtectedRoute>
                            <Dashboard />
                        </ProtectedRoute>
                    }
                >
                    <Route path="providers" element={<ProvidersList />} />
                    <Route path="providers/:id" element={<ProviderDetails />} />

                    <Route path="/dashboard/events" element={<AppEvents/>} />
                    <Route path="/dashboard/workflow" element={<Workflow/>} />
                    <Route path="/dashboard/users" element={<ConnectedAccounts/>} />
                    <Route path="/dashboard/tasks" element={<TaskHistory/>} />
                    <Route path="/dashboard/docs" element={<></>} />
                    <Route path="/dashboard/settings" element={<></>} />
                    <Route path="/dashboard/help" element={<></>} />
                </Route>

                <Route path="/sdk-test" element={<SdkTestPage />} />
            </Routes>
        </StrictMode>
    </BrowserRouter>
);
