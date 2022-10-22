import { Box, Button, Divider, Input, Paper, Typography } from "@mui/material";
import AddIcon from '@mui/icons-material/Add'; 
import React from "react";
import { CreateDeckDialog } from "./components/create-deck-dialog";
import { ApiClientMethod, useApi } from "../../../hooks/use-api";
import { IDeck } from "../../../types/deck";
import { useSnackbar } from "notistack";
import { Link } from "react-router-dom";
import { LoadingSmall } from "../../../components/loading-small";
import { EditDeckDialog } from "./components/edit-deck-dialog";
import { LoadingBackdrop } from "../../../components/loading-backdrop";
import { useCalculateHeight } from "../../../hooks/use-calculate-height";

export const ManageDecks: React.FC = () => {
  // snackbar
  const { enqueueSnackbar } = useSnackbar();

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
  const [deck, setDeck] = React.useState<IDeck | null>(null);

  const handleClickOpenEditDialog = (deck: IDeck) => {
    setDeck(deck);
    setEditDialogOpen(true);
  };

  const handleCloseEditDialog = () => {
    setEditDialogOpen(false);
  };

  // fetch decks
  const [fetch, { data, setData, loading }] = useApi<void, Array<IDeck>>({
    method: ApiClientMethod.GET,
    url: "/api/decks/premade",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully fetched decks", {variant: "success"});
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  // delete decks
  const [deleteDeck, {loading: deleteDeckLoading}] = useApi<void, void>({
    method: ApiClientMethod.DELETE,
    url: "/api/decks",
    onSuccess: (data) => {
      enqueueSnackbar("Successfully deleted deck", {variant: "success"});
      console.log(data);
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  const deleteHandler = async (deckId: number) => {
    await deleteDeck({urlAddition: deckId.toString()});
    setData((prev) => prev.filter(deck => deck.id !== deckId));
  }

  React.useEffect(() => {
    fetch({});
  }, [fetch]);

  // search
  const [search, setSearch] = React.useState<string>("");

  const calculatedHeight = useCalculateHeight()

  return (
    <>
    <CreateDeckDialog open={dialogOpen} handleClose={handleCloseDialog} setDeckList={setData}/>
    <EditDeckDialog open={editDialogOpen} handleClose={handleCloseEditDialog} deck={deck!} setDeckList={setData}/>
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
          <Typography component="h3" variant="h5" sx={{ fontWeight: 600 }}>Decks</Typography>
          <Button variant="contained" startIcon={<AddIcon />} onClick={handleClickOpenDialog}>Create</Button>
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
            data && data.filter((deck) => deck.name.includes(search) || deck.id.toString() === search).map((deck) => (
              <Paper key={deck.id} variant="outlined" sx={{
                padding: 2,
                minHeight: 75,
                display: "flex",
                flexWrap: "nowrap",
                justifyContent: "space-between",
                alignItems: "center",
                marginY: 2,
                textDecoration: "none"
              }}>
                <Typography variant="h5">{`${deck.name} (${deck.amount} ${deck.amount === 1 ? "card" : "cards"})`}</Typography>
                <Box sx={{
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                  flexWrap: "wrap"
                }}>
                  <Button component={Link} to={"/admin/decks/" + deck.id} color="success">Show</Button>
                  <Button color="info" onClick={() => {
                    handleClickOpenEditDialog(deck);
                  }}>Edit</Button>
                  <Button color="error" onClick={() => deleteHandler(deck.id)}>Delete</Button>
                </Box>
              </Paper>
            ))
          }
          {loading && <LoadingSmall />}
        </Box>
      </Box>
    }
    {deleteDeckLoading && <LoadingBackdrop />}
    </>
  );
}
