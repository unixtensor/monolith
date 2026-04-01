import { useEffect, useState } from "react";
import { ServerAliveContext } from "./context";

export default function ServerAlive({
	children,
}: {
	children: React.ReactNode;
}) {
	const [alive, setAlive] = useState<boolean | null>(null);

	useEffect(() => {
		fetch("/api/v1/connected").then((r) => setAlive(r.ok));
	}, []);
	return (
		<ServerAliveContext.Provider value={alive}>
			{children}
		</ServerAliveContext.Provider>
	);
}
