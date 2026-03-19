import { Route, Routes } from "react-router";
import Index from "./app/init";

function App() {
	return (
		<Routes>
			<Route index element={<Index />} />
		</Routes>
	);
}

export default App;
