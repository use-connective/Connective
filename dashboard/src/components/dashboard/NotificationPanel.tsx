import {
    Popover,
    PopoverTrigger,
    PopoverContent,
} from "@/components/ui/popover";
import { Bell } from "lucide-react";

export function NotificationDropdown() {
    return (
        <Popover>
            <PopoverTrigger asChild>
                <Bell className="w-5 h-5 text-white cursor-pointer" />
            </PopoverTrigger>

            <PopoverContent className="w-64">
                <h3 className="text-sm font-semibold mb-2">Notifications</h3>

                <div className="space-y-2">
                    <div className="p-2 border rounded">User connected</div>
                    <div className="p-2 border rounded">Task completed</div>
                </div>
            </PopoverContent>
        </Popover>
    );
}
