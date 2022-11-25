import { Cards } from "@/features/admin/cards";
import * as React from "react";
import { useNavigate, useParams } from "react-router-dom";

export const AdminCardsPage = () => {
	const { deckId } = useParams();
	const navigate = useNavigate();
	React.useEffect(() => {
		if (!/\d+/.test(deckId!)) {
			navigate(-1);
		}}, [deckId, navigate]);

	return (
		<Cards deckID={+(deckId as string) as number} />
	);
}


