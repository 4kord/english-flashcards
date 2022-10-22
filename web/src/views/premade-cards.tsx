import { Box } from "@mui/system";
import { useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { PremadeCards } from "../features/user/premade-cards";

export const PremadeCardsView = () => {
  const {deckId} = useParams();
  const navigate = useNavigate();
  useEffect(() => {
    if (!/\d+/.test(deckId!)) {
      navigate(-1);
    }}, [deckId, navigate]);

  return (
    <Box>
      <PremadeCards deckId={+deckId!}/>
    </Box>
  );
}
