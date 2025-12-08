import {Input} from "@/components/ui/input.tsx";
import {Search} from "lucide-react";
import {useEffect, useState} from "react";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow,} from "@/components/ui/table"
import type {ConnectedAccount} from "@/api/dto/integration.ts";
import {getConnectedAccounts} from "@/api/integration.ts";
import {toast} from "react-hot-toast";
import {useProjectStore} from "@/store/projectStore.ts";
import { FiInbox } from "react-icons/fi";


export default function ConnectedAccounts() {
    const [searchText, setSearchText] = useState('')
    const [connectedAccounts, setConnectedAccounts] = useState<ConnectedAccount[]>()
    const project = useProjectStore(state => state.selectedProject)


    async function fetchConnectedAccounts() {
        if (!project) {
            toast.error("No project selected.")
            return
        }

        try {
            const accounts = await getConnectedAccounts(project?.id, searchText)
            setConnectedAccounts(accounts.data)
        } catch (err) {
            const message = err instanceof Error ? err.message : "Something went wrong"
            toast.error(message)
        }
    }

    useEffect(() => {
        if (project) {
            fetchConnectedAccounts()
        }
    }, [project]);


    return (
        <div className="flex flex-col space-y-6 h-full">

            {/* Search Row */}
            <div className="flex space-x-4 w-1/3">
                <Input
                    className="h-12 bg-white"
                    placeholder="User Id"
                    value={searchText}
                    onChange={(e) => setSearchText(e.target.value)}
                />
                <div className="h-12 w-14 bg-slate-700 rounded-md flex items-center justify-center cursor-pointer" onClick={() => {
                    fetchConnectedAccounts()
                }}>
                    <Search className="text-white"/>
                </div>
            </div>

            <div className="flex-1 bg-white border border-gray-200 rounded-2xl p-6 overflow-auto">

                {
                    connectedAccounts === undefined || connectedAccounts === null || connectedAccounts.length === 0 ?
                        <div className="flex flex-col space-y-4 items-center justify-center h-full">
                            <FiInbox className="text-5xl text-slate-800" />
                            <p
                            className="text-xl font-medium text-slate-800">Not Connected Account Found</p>

                            <p className="text-center">There are currently no connected integrations for this project. <br/>
                                Once they complete an integration, their accounts will be displayed here.</p>
                        </div> :
                        <Table className="w-full table-fixed">
                            <TableHeader>
                                <TableRow>
                                    <TableHead className="w-[30%]">User ID</TableHead>
                                    <TableHead className="w-[55%]">Integrations Enabled</TableHead>
                                    <TableHead className="w-[15%]">Date Created</TableHead>
                                </TableRow>
                            </TableHeader>

                            <TableBody>
                                {
                                    connectedAccounts.map((user) => (
                                        <TableRow key={user.user_id}>
                                            <TableCell className="w-[30%] font-medium">{user.user_id}</TableCell>
                                            <TableCell className="w-[55%]">{user.integrations_enabled}</TableCell>
                                            <TableCell className="w-[15%]">{user.displayable_date}</TableCell>
                                        </TableRow>
                                    ))
                                }
                            </TableBody>
                        </Table>
                }

            </div>

        </div>
    )
}
