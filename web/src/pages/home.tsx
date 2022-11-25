import { useAuth } from "@/hooks/use-auth";

export const HomePage = () => {
	const { auth } = useAuth();

	return (
		<h1>{JSON.stringify(auth)}</h1>
	);
}
