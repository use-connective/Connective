import {Search} from "lucide-react";
import {Input} from "../ui/input";
import {useEffect, useState} from "react";
import type {ProviderDisplayable} from "@/api/dto/integration";
import {getAllCategories, getAllProviders} from "@/api/integration";
import toast from "react-hot-toast";
import {Link} from "react-router";

export default function ProvidersList() {
    const [providers, setProviders] = useState<ProviderDisplayable[]>([]);
    const [searchText, setSearchText] = useState('')
    const [category, setCategory] = useState('All')
    const [categories, setCategories] = useState<string[]>()

    const fetchProviders = async () => {
        try {
            const resp = await getAllProviders(searchText, category);
            setProviders(resp);
        } catch (err) {
            toast.error("Failed to load dashboard.");
        }
    };

    const fetchCategories = async () => {
        try {
            const resp = await getAllCategories()
            setCategories(resp.data)
        } catch (err) {
            toast.error("Failed to load categories.");
        }
    };

    useEffect(() => {
        fetchProviders();
    }, [category]);

    useEffect(() => {
        fetchCategories();
    }, []);

    return (
        <div className="flex space-x-6 h-full">
            {/* LEFT: Categories */}
            <div className="flex-1 flex flex-col space-y-4">
                {/* Search + Open Issue */}
                <div className="flex justify-between">
                    <div className="flex space-x-4 w-1/2">
                        <Input
                            className="h-12 bg-white"
                            placeholder="Integration Name"
                            value={searchText}
                            onChange={(e) => {
                                setSearchText(e.target.value)
                            }}
                        />
                        <div
                            className="h-12 w-14 bg-slate-700 rounded-md flex items-center justify-center cursor-pointer">
                            <Search className="text-white" onClick={() => {
                                if (searchText === '') {
                                    toast.error("Enter integration name first.")
                                } else {
                                    toast.loading('Searching')
                                    fetchProviders()
                                    toast.dismiss()
                                }
                            }}/>
                        </div>
                    </div>

                    <div className="h-12 flex items-center bg-white border border-gray-200 rounded-md px-4 text-xs">
                        ✨ Don’t see the integration you're looking for?{" "}
                        <span className="text-[#4268FB] font-bold ml-1 cursor-pointer">
                            Open an Issue.
                        </span>
                    </div>
                </div>

                {/* Main Integration Grid Card */}
                <div className="flex-1 bg-white border border-gray-200 rounded-2xl p-6 overflow-auto">
                    <h2 className="text-md font-semibold text-gray-800 mb-6">
                        {category}
                    </h2>

                    <div className="grid grid-cols-4 gap-6">
                        {providers === undefined || providers.length === 0 ? (
                            <div
                                className="w-full col-span-6 flex flex-col items-center justify-center py-20 text-gray-700">
                                <Search className="w-12 h-12 mb-4 opacity-80"/>
                                <p className="text-lg font-medium">No integrations found</p>
                                <p className="text-sm text-gray-400 mt-1">
                                    Try changing the category or search term.
                                </p>
                            </div>
                        ) : (
                            providers.map((provider) => (
                                <Link
                                    to={`/dashboard/providers/${provider.id}`}
                                    key={provider.id}
                                    className="border rounded-xl overflow-hidden hover:border-slate-400"
                                >
                                    <div className="flex flex-col items-center justify-center py-10 bg-white">
                                        <img
                                            src={provider.image_url}
                                            alt={provider.display_name}
                                            className="w-12 h-12"
                                        />
                                    </div>
                                    <div className="border-t py-3 bg-white text-center">
                                        <p className="text-sm font-medium text-gray-900">
                                            {provider.display_name}
                                        </p>
                                    </div>
                                </Link>
                            ))
                        )}

                    </div>

                </div>
            </div>


            {/* RIGHT SIDE */}
            <div className="w-72 bg-white border border-gray-200 rounded-2xl p-6">
                <h2 className="text-sm font-semibold text-gray-700 mb-6">
                    CATEGORIES
                </h2>


                <div className="flex flex-col gap-4 text-sm">
                    {categories?.map((cat) => (
                        <div
                            key={cat}
                            className={`flex justify-between items-center cursor-pointer ${
                                category === cat ? "text-slate-700 font-bold" : "text-gray-700 hover:text-blue-600"
                            }`}
                            onClick={() => setCategory(cat)}
                        >
                            <span>{cat}</span>
                        </div>
                    ))}
                </div>
            </div>

        </div>
    );
}
