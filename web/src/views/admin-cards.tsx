import React from "react";
import { Box } from "@mui/system";
import { useNavigate, useParams } from "react-router-dom";
import { ManageCards } from "../features/admin/manage-cards";

export const AdminCardsView: React.FC = () => {
  const {deckId} = useParams();
  const navigate = useNavigate();
  React.useEffect(() => {
    if (!/\d+/.test(deckId!)) {
      navigate(-1);
    }}, [deckId, navigate]);

  return (
    <Box>
      <ManageCards deckId={deckId!} />
    </Box>
  )
}
