import { Box, Button, Divider, Input, Paper, Typography } from "@mui/material";
import AddIcon from '@mui/icons-material/Add'; 
import React from "react";
import { ICard } from "../../../types/card";
import { ApiClientMethod, useApi } from "../../../hooks/use-api";
import { useSnackbar } from "notistack";
import { AddCardDialog } from "./components/add-card-dialog";
import { LoadingSmall } from "../../../components/loading-small";
import { LoadingBackdrop } from "../../../components/loading-backdrop";
import { PreviewCardDialog } from "./components/preview-card-dialog/component";
import { EditCardDialog } from "./components/edit-card-dialog";
import { useCalculateHeight } from "../../../hooks/use-calculate-height";

interface IProps {
  deckId: string;
}

export const ManageCards: React.FC<IProps> = ({deckId}) => {
  // snackbar
  const { enqueueSnackbar } = useSnackbar();

  // search
  const [search, setSearch] = React.useState<string>("");

  // get cards
  const [fetch, {data, setData, loading}] = useApi<void, Array<ICard>>({
    method: ApiClientMethod.GET,
    url: "/api/decks",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully fetched cards", {variant:"success"});
    },
    onFail: (error) => {
      console.error(error);
      enqueueSnackbar(error.errorCode, {variant:"error"});
    }
  });

  // open and close dialog
  const [dialogOpen, setDialogOpen] = React.useState(false);

  const handleClickOpenDialog = () => {
    setDialogOpen(true);
  };

  const handleCloseDialog = () => {
    setDialogOpen(false);
  };

  // open and close edit dialog
  const [editDialogOpen, setEditDialogOpen] = React.useState(false);
  const [editCard, setEditCard] = React.useState<ICard | null>(null);

  const handleClickOpenEditDialog = (card: ICard) => {
    setEditDialogOpen(true);
    setEditCard(card);
  };

  const handleCloseEditDialog = () => {
    setEditDialogOpen(false);
  };

  React.useEffect(() => {
    fetch({
      urlAddition: deckId
    });
  }, [fetch, deckId]);

  // delete decks
  const [deleteCard, {loading: deleteCardLoading}] = useApi<void, void>({
    method: ApiClientMethod.DELETE,
    url: "/api/cards",
    onSuccess: (data) => {
      enqueueSnackbar("Successfully deleted card", {variant: "success"});
      console.log(data);
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  const deleteHandler = async (cardId: number) => {
    await deleteCard({urlAddition: cardId.toString()});
    setData((prev) => prev.filter(card => card.id !== cardId));
  }

  // open and close preview dialog
  const [previewDialogOpen, setPreviewDialogOpen] = React.useState<boolean>(false);
  const [previewCard, setPreviewCard] = React.useState<ICard | null>(null);

  const handleClickOpenPreviewDialog = (card: ICard) => {
    setPreviewDialogOpen(true);
    setPreviewCard(card);
  };

  const handleClosePreviewDialog = () => {
    setPreviewDialogOpen(false);
  };

  const calculatedHeight = useCalculateHeight()

  return (
    <>
    <AddCardDialog open={dialogOpen} handleClose={handleCloseDialog} setCardList={setData} deckId={deckId} />
    <EditCardDialog open={editDialogOpen} handleClose={handleCloseEditDialog} card={editCard} setCardList={setData} />
    <PreviewCardDialog open={previewDialogOpen} handleClose={handleClosePreviewDialog} card={previewCard} />
    {calculatedHeight &&
      <Box sx={{
        overflowY: "scroll",
        height: calculatedHeight
      }}>
        <Box sx={{
          padding: 2,
          display: "flex",
          justifyContent: "space-between",
        }}>
          <Typography component="h3" variant="h5" sx={{ fontWeight: 600 }}>Cards</Typography>
          <Button variant="contained" startIcon={<AddIcon />} onClick={handleClickOpenDialog}>Add</Button>
        </Box>
        <Divider />
        <Box>
          <Input placeholder="Search" fullWidth onChange={(e: React.ChangeEvent<HTMLInputElement>) => setSearch(e.target.value)} sx={{
            padding: 2
          }}/>
        </Box>
        <Box sx={{
          padding: 2,
        }}>
          {
            data && data.filter((card) => card.english.includes(search) || card.russian.includes(search) || card.id.toString() === search).map((card) => (
              <Paper key={card.id} variant="outlined" sx={{
                padding: 2,
                minHeight: 75,
                display: "flex",
                flexWrap: "nowrap",
                justifyContent: "space-between",
                alignItems: "center",
                marginY: 2
              }}>
                <Typography variant="h5">{card.english}</Typography>
                <Typography variant="h5">{card.russian}</Typography>
                <Typography variant="h5">{card.association}</Typography>
                <Box sx={{
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                  flexWrap: "wrap"
                }}>
                  <Button color="success" onClick={() => handleClickOpenPreviewDialog(card)}>Preview</Button>
                  <Button color="info" onClick={() => handleClickOpenEditDialog(card)}>Edit</Button>
                  <Button color="error" onClick={() => deleteHandler(card.id)}>Delete</Button>
                </Box>
              </Paper>
            ))
          }
          {loading && <LoadingSmall />}
        </Box>
      </Box>
    }
    {deleteCardLoading && <LoadingBackdrop />}
    </>
  );
}
