import { lazy } from "react";
import { Navigate, Route, Routes } from "react-router";
import Auth from "./app/init";
import Login from "./app/login/init";

const Dashboard = lazy(() => import("./app/dashboard/init"));
const Servers = lazy(() => import("./app/servers/init"));

function App() {
	return (
		<Routes>
			<Route path="/login" element={<Login />} />
			<Route path="*" element={<Navigate to="/servers" replace />} />
			<Route element={<Auth />}>
				<Route element={<Dashboard />}>
					<Route path="/servers" index element={<Servers />} />
					<Route path="/:placeId/" />
					<Route path="/:placeId/:jobId" />
				</Route>
			</Route>
		</Routes>
	);
}

export default App;
