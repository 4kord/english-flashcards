import AddIcon from '@mui/icons-material/Add';
import { LoadingSmall } from "@/components/loading-small";
import { Box, Button, Divider, Paper, Typography } from "@mui/material";
import { useSnackbar } from 'notistack';
import { ApiClientMethod, useApi } from '@/hooks/use-api';
import { useCalculateHeight } from '@/hooks/use-calculate-height';
import { useEffect, useState } from 'react';
import { Card } from '@/models/card';
import { AddCardModal } from './components/add-card-modal';
import { CardPreviewModal } from './components/card-preview-modal';

const CardsBody: React.FC<{cards: Array<Card>, setCards: React.Dispatch<React.SetStateAction<Card[]>>, setPreviewCard: React.Dispatch<React.SetStateAction<Card | undefined>>, setOpenPreview: React.Dispatch<React.SetStateAction<boolean>>, loading: boolean}> = ({cards, setCards, setPreviewCard, setOpenPreview, loading}) => {
	const snackbar = useSnackbar();

	const [deleteCard] = useApi<void, void>({
		method: ApiClientMethod.DELETE,
		url: `/cards/$`,
		onSuccess: (data) => {
			console.log(data);
			snackbar.enqueueSnackbar("Successfully deleted card", { variant: "success" });
		},
		onFail: (error) => {
			console.log(error);
			snackbar.enqueueSnackbar(error.error.code, { variant: "error" });
		},
	});

	const handlePreview = (card: Card) => {
		setOpenPreview(true);
		setPreviewCard(card);
	}

	const handleEdit = () => {

	}

	const handleDelete = async (cardID: number) => {
		await deleteCard({
			params: [cardID.toString()]
		});

		setCards((prev) => prev.filter(card => card.id !== cardID));
	}

	return (
		<>
		<Box sx={{
			paddingX: 2,
		}}>
			{
				cards && cards.map((card) => (
					<Paper key={card.id} variant="outlined" sx={{
						padding: 2,
						minHeight: 75,
						display: "flex",
						flexWrap: "nowrap",
						justifyContent: "space-between",
						alignItems: "center",
						marginY: 2
					}}>
						<Typography variant="h5">{`${card.english} - ${card.russian} - ${card.association}`}</Typography>
						<Box>
							<Button color="success" onClick={() => handlePreview(card)}>Preview</Button>
							<Button color="info" onClick={handleEdit}>Edit</Button>
							<Button color="error" onClick={() => handleDelete(card.id)}>Delete</Button>
						</Box>
					</Paper>
				))
			}
		</Box>
		{loading && <LoadingSmall />}
		</>
	);
}

const CardsHeader: React.FC<{handleAddCard: () => void}> = ({handleAddCard}) => {
	return (
		<Box sx={{
			padding: 2,
			display: "flex",
			justifyContent: "space-between",
		}}>
			<Typography component="h3" variant="h5" sx={{ fontWeight: 600 }}>Cards</Typography>
			<Button variant="contained" startIcon={<AddIcon />} onClick={handleAddCard}>Add</Button>
		</Box>
	);
}

export const Cards: React.FC<{deckID: number}> = ({ deckID }) => {
	const snackbar = useSnackbar();
	const calculatedHeight = useCalculateHeight();

	const [getCards, { data: cards, setData: setCards, loading }] = useApi<void, Array<Card>>({
		method: ApiClientMethod.GET,
		url: `/decks/${deckID}/cards`,
		onSuccess: (data) => {
			console.log(data);
			snackbar.enqueueSnackbar("Successfully fetched cards", { variant: "success" });
		},
		onFail: (error) => {
			console.log(error);
			snackbar.enqueueSnackbar(error.error.code, { variant: "error" });
		}
	});

	useEffect(() => {
		getCards({});
	}, []);

	// Add Deck Modal
	const [open, setOpen] = useState(false);

	// Card Prview Modal
	const [openPreview, setOpenPreview] = useState(false);
	const [previewCard, setPreviewCard] = useState<Card | undefined>(undefined);

	return (
		<>
		<AddCardModal open={open} handleClose={() => setOpen(false)} setCardList={setCards} deckID={deckID} />
		<CardPreviewModal open={openPreview} handleClose={() => setOpenPreview(false)} card={previewCard} />
    {calculatedHeight &&
      <Box sx={{
        overflowY: "scroll",
        height: calculatedHeight
      }}>
				<CardsHeader handleAddCard={() => setOpen(true)}/>
        <Divider />
				<CardsBody cards={cards} setCards={setCards} setPreviewCard={setPreviewCard} setOpenPreview={setOpenPreview} loading={loading} />
      </Box>
    }
		</>
	);
}
