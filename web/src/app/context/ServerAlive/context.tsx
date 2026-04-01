import { createContext, useContext } from "react";

export const ServerAliveContext = createContext<boolean | null>(null);

export default function useServerAlive() {
	return useContext(ServerAliveContext);
}
