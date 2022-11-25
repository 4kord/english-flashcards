import AddIcon from '@mui/icons-material/Add';
import { LoadingSmall } from "@/components/loading-small";
import { Box, Button, Divider, Paper, Typography } from "@mui/material";
import { Deck } from '@/models/deck';
import { useSnackbar } from 'notistack';
import { ApiClientMethod, useApi } from '@/hooks/use-api';
import { useCalculateHeight } from '@/hooks/use-calculate-height';
import { AddDeckModel } from './components/add-deck-modal';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

const DecksBody: React.FC<{decks: Array<Deck>, loading: boolean}> = ({decks, loading}) => {
	const navigate = useNavigate();

	return (
		<>
		<Box sx={{
			paddingX: 2,
		}}>
			{
				decks && decks.map((deck) => (
					<Paper key={deck.id} variant="outlined" sx={{
						padding: 2,
						minHeight: 75,
						display: "flex",
						flexWrap: "nowrap",
						justifyContent: "space-between",
						alignItems: "center",
						marginY: 2
					}}>
						<Typography variant="h5">{`${deck.name} (${deck.amount} cards)`}</Typography>
						<Box>
							<Button color="success" onClick={() => navigate(`/admin/decks/${deck.id}`)}>View</Button>
							<Button color="info" onClick={() => {}}>Edit</Button>
							<Button color="error" onClick={() => {}}>Delete</Button>
						</Box>
					</Paper>
				))
			}
		</Box>
		{loading && <LoadingSmall />}
		</>
	);
}

const DecksHeader: React.FC<{handleAddDeck: () => void}> = ({handleAddDeck}) => {
	return (
		<Box sx={{
			padding: 2,
			display: "flex",
			justifyContent: "space-between",
		}}>
			<Typography component="h3" variant="h5" sx={{ fontWeight: 600 }}>Decks</Typography>
			<Button variant="contained" startIcon={<AddIcon />} onClick={handleAddDeck}>Add</Button>
		</Box>
	);
}

export const Decks: React.FC = () => {
	const snackbar = useSnackbar();
	const calculatedHeight = useCalculateHeight()

	const [getDecks, { data: decks, setData: setDecks, loading }] = useApi<void, Array<Deck>>({
		method: ApiClientMethod.GET,
		url: "/decks/premade",
		onSuccess: (data) => {
			console.log(data);
			snackbar.enqueueSnackbar("Successfully fetched decks", { variant: "success" })
		},
		onFail: (error) => {
			console.log(error);
			snackbar.enqueueSnackbar(error.error.code, { variant: "error" })
		}
	});

	useEffect(() => {
		getDecks();
	}, []);

	// Add Deck Modal
	const [open, setOpen] = useState(false);

	return (
		<>
		<AddDeckModel open={open} handleClose={() => setOpen(false)} setDeckList={setDecks} />
    {calculatedHeight &&
      <Box sx={{
        overflowY: "scroll",
        height: calculatedHeight
      }}>
				<DecksHeader handleAddDeck={() => setOpen(true)}/>
        <Divider />
				<DecksBody decks={decks} loading={loading} />
      </Box>
    }
		</>
	);
}
