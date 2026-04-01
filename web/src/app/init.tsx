import { Card } from "@/components/ui/card";
import { CircleX, LoaderPinwheel } from "lucide-react";
import { Outlet } from "react-router";
import useServerAlive from "./context/ServerAlive/context";
import { useTitle } from "./hooks/useTitle";

function Loading() {
	return (
		<main className="flex justify-center items-center h-screen">
			<LoaderPinwheel className="animate-spin-pulse text-[#b3b3b3] w-10 h-10" />
		</main>
	);
}

function Error() {
	useTitle("Server Error");

	return (
		<main className="w-screen h-screen flex justify-center items-center">
			<Card className="flex flex-col content-center items-center p-8 gap-5 bg-[#140003]">
				<CircleX />
				<p>Server is unreachable</p>
			</Card>
		</main>
	);
}

export default function GetConnected() {
	const s_alive = useServerAlive();

	if (s_alive === null) return <Loading />;
	if (s_alive === false) return <Error />;
	// if (s_alive === Connected.False) return <Navigate to="/login" replace />;
	return <Outlet />;
}
