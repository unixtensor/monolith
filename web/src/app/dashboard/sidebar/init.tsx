import { Button } from "@/components/ui/button";
import {
	Sidebar as ShadSidebar,
	SidebarContent,
	SidebarFooter,
	SidebarGroup,
	SidebarGroupLabel,
	SidebarHeader,
} from "@/components/ui/sidebar";
import { GitForkIcon, ServerIcon } from "lucide-react";
import Logout from "./logout";

function ButtonsGroup() {
	return (
		<SidebarGroup>
			<Button className="justify-baseline">
				<ServerIcon /> Servers
			</Button>
			<Button className="justify-baseline bg-sidebar">
				<GitForkIcon /> Graph
			</Button>
		</SidebarGroup>
	);
}

function SettingsGroup() {
	return (
		<SidebarGroup>
			<SidebarGroupLabel>Settings</SidebarGroupLabel>
		</SidebarGroup>
	);
}

export default function Sidebar() {
	return (
		<ShadSidebar className="border-none">
			<SidebarHeader>
				<h1 className="text-center">Monolith</h1>
			</SidebarHeader>
			<SidebarContent>
				<ButtonsGroup />
				<SettingsGroup />
			</SidebarContent>
			<SidebarFooter>
				<Logout />
			</SidebarFooter>
		</ShadSidebar>
	);
}
