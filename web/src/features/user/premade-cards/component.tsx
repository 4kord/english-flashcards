import { Button, FormControl, InputLabel, MenuItem, Select, SelectChangeEvent } from "@mui/material";
import { Box } from "@mui/system";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { useSnackbar } from "notistack";
import React from "react";
import { LoadingSmall } from "../../../components/loading-small";
import { NoContent } from "../../../components/nocontent";
import { ApiClientMethod, useApi } from "../../../hooks/use-api";
import { useCalculateHeight } from "../../../hooks/use-calculate-height";
import { ICard } from "../../../types/card";
import { IDeck } from "../../../types/deck";
import { IRequest } from "./types";

interface IProps {
  deckId: number
}

const columns: GridColDef[] = [
  { field: 'english', headerName: 'English', flex: 1},
  { field: 'russian', headerName: 'Russian', flex: 1},
  { field: 'createdAt', headerName: 'Created At', type: "dateTime", flex: 1, valueFormatter: params => new Date(params.value).toUTCString()}
];

export const PremadeCards: React.FC<IProps> = ({deckId}) => {
  // calculated height
  const calculatedHeight = useCalculateHeight()

  // snackbar
  const {enqueueSnackbar} = useSnackbar();

  // fetch cards
  const [fetch, { data, loading }] = useApi<void, Array<ICard>>({
    method: ApiClientMethod.GET,
    url: `/api/decks/${deckId}`,
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully fetched cards", {variant: "success"});
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  // fetch user's decks
  const [fetchUserDecks, {data: userDecks}] = useApi<void, Array<IDeck>>({
    method: ApiClientMethod.GET,
    url: `/api/decks`,
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully fetched user's decks", {variant: "success"});
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  React.useEffect(() => {
    const initFetch = async () => {
      await fetch({});
      await fetchUserDecks({});
    }

    initFetch();
  }, [fetch, fetchUserDecks, deckId])

  const [selectedIds, setSelectedIds] = React.useState<Array<number>>([]);
  const [destination, setDestination] = React.useState<string>("");

  const handleChange = (event: SelectChangeEvent) => {
    setDestination(event.target.value as string);
  };

  // copy cards
  const [copyCards] = useApi<IRequest, void>({
    method: ApiClientMethod.POST,
    url: `/api/decks/${Number(destination)}/copy`,
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully copied cards", {variant: "success"});
    },
    onFail: (error) => {
      console.log(error)
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  const handleCopy = async () => {
    console.log(selectedIds);
    await copyCards({
      data: {
        ids: selectedIds
      }
    });
    setSelectedIds([]);
  }


  return (
    <Box sx={{
      overflowY: "scroll",
    }}>
      { data && !loading && calculatedHeight && <Box sx={{
          paddingTop: 2,
          paddingX: 2,
          height: calculatedHeight
        }}>
          <div style={{ height: "100%", width: "100%", display: "flex", flexDirection: "column"}}>
            <DataGrid
              rows={data}
              columns={columns}
              disableSelectionOnClick
              checkboxSelection
              onSelectionModelChange={ids => {
                setSelectedIds(ids as Array<number>);
              }}
              selectionModel={selectedIds}
            />
            <Box sx={{display: "flex", gap: 2, marginY: 2}}>
              <FormControl sx={{flex: 1}}>
                <InputLabel id="select-deck-label">Destination Deck</InputLabel>
                <Select
                  labelId="select-deck-label"
                  label="Destination Deck"
                  value={destination}
                  onChange={handleChange}
                >
                  { userDecks && userDecks.map(deck => (
                    <MenuItem key={deck.id} value={deck.id}>{deck.name}</MenuItem>
                  ))
                  }
                </Select>
              </FormControl>
              <Button variant="contained" sx={{flex: 1}} disabled={selectedIds.length === 0 || !destination ? true : false} onClick={handleCopy}>Add Selected Cards</Button>
            </Box>
          </div>
        </Box>
      }
      { !data && !loading &&
        <NoContent />
      }
      {loading && <LoadingSmall />}
    </Box>
  );
}
