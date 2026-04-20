import { Button } from "@/components/ui/button";
import {
	CircleXIcon,
	ExternalLinkIcon,
	HardDriveIcon,
	LoaderCircleIcon,
	TagIcon,
	UsersIcon,
} from "lucide-react";
import { Link } from "react-router";
import { useGames } from "../dashboard/games-servers";
import SearchProvider, { useSearch } from "./search";

function NoResult({ children }: { children: string }) {
	return (
		<div className="flex flex-col gap-5 justify-center items-center h-100">
			<CircleXIcon className="h-10 w-10" />
			<h1>{children}</h1>
		</div>
	);
}

function Loading() {
	return (
		<div className="flex justify-center items-center h-100">
			<LoaderCircleIcon className="animate-spin h-10 w-10" />
		</div>
	);
}

function DisplayGames() {
	const games = useGames();
	const search = useSearch();

	if (games.isLoading) return <Loading />;
	if (games.data.length === 0)
		return <NoResult>No games are running</NoResult>;

	const filtered = games.data.filter(
		(game) =>
			game.Name.toLowerCase().includes(search.searchTerm.toLowerCase()) ||
			String(game.PlaceId).includes(search.searchTerm),
	);
	if (filtered.length === 0)
		return (
			<NoResult>{`No game with name nor id "${search.searchTerm}" found`}</NoResult>
		);

	return filtered.map((game) => (
		<Link to={`/${game.PlaceId}`} key={game.JobId}>
			<Button className="w-full bg-secondary hover:bg-primary flex justify-between gap-5 min-h-22 p-5">
				<div className="flex items-center gap-5">
					<img
						src={
							game.JobId !== "studio"
								? "https://upload.wikimedia.org/wikipedia/commons/2/29/Roblox_Icon_Transparent.png"
								: "https://upload.wikimedia.org/wikipedia/commons/e/ee/Roblox_Studio_icon_2025.svg"
						}
						className="w-15 h-15"
					/>
					<div className="flex flex-col gap-2 text-left">
						<div className="flex items-center gap-3">
							<h1 className="font-bold text-lg">{game.Name}</h1>
							<div className="bg-green-500 w-3 h-3 rounded-full"></div>
						</div>
						<div className="flex gap-3 opacity-60 [&>div]:flex [&>div]:items-center [&>div]:gap-1">
							<div>
								<TagIcon />
								<h2>{game.PlaceId}</h2>
							</div>
							<div>
								<UsersIcon />
								<h2>{game.MaxPlayers}</h2>
							</div>
						</div>
					</div>
				</div>
				<ExternalLinkIcon />
			</Button>
		</Link>
	));
}

export default function Games() {
	return (
		<>
			<header className="flex gap-3 mt-3 mb-5">
				<HardDriveIcon />
				<div>Active Games</div>
			</header>
			<div className="flex flex-col gap-5 mt-3">
				<SearchProvider queryKey={["games-servers"]}>
					<DisplayGames />
				</SearchProvider>
			</div>
		</>
	);
}
