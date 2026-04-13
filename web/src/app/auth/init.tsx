import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { Context } from "./context";
import { logged_in, need_login } from "./utils";

export interface AuthContext {
	guest?: boolean;
	isLoading: boolean;
	error: Error | null;
}

export default function Auth({ children }: { children: React.ReactNode }) {
	const { data, isLoading, error } = useQuery({
		queryKey: ["auth"],
		queryFn: () =>
			axios
				.get("/api/v1", {
					validateStatus: (s) => need_login(s) || logged_in(s),
				})
				.then((r) => need_login(r.status)),
	});
	return (
		<Context.Provider value={{ guest: data, isLoading, error }}>
			{children}
		</Context.Provider>
	);
}
