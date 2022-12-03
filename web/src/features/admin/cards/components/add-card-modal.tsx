import { Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, Switch, TextField, Typography } from "@mui/material";
import { Box } from "@mui/system";
import CloseIcon from '@mui/icons-material/Close';
import React, { Dispatch, FormEvent, SetStateAction, useState } from "react";
import { ApiClientMethod, useApi } from "@/hooks/use-api";
import { LoadingBackdrop } from "@/components/loading-backdrop";
import { useSnackbar } from "notistack";
import { AddCard, Card } from "@/models/card";
import { api } from "@/lib/axios";

interface Props {
  open: boolean;
  handleClose: () => void;
  setCardList: Dispatch<SetStateAction<Card[]>>;
	deckID: number;
}

export const AddCardModal: React.FC<Props> = ({ open, handleClose, setCardList, deckID }) => {
  const snackbar = useSnackbar();

	const [formData, setFormData] = useState<AddCard>({
		english: "",
		russian: "",
		association: "",
		example: "",
		transcription: "",
		image: undefined,
		image_url: "",
		audio: undefined,
		audio_url: ""
	})

	const [audioChecked, setAudioChecked] = useState<boolean>(false);

  const [addCard, { loading }] = useApi<AddCard, Card>({
    method: ApiClientMethod.POST,
    url: `/decks/${deckID}/cards`,
		headers: {"Content-Type": "multipart/form-data"},
    onSuccess: (data) => {
      console.log(data);
      snackbar.enqueueSnackbar("Successfully added deck", {variant: "success"});
      setCardList(prev => {
        if (prev) {
          return [...prev, data];
        } else {
          return [data];
        }
      });
    },
    onFail: (error) => {
      console.log(error);
      snackbar.enqueueSnackbar(error.error.code, {variant: "error"});
    }
  });

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    console.log(formData);

    await addCard({
			data: formData
		})

    setFormData({
			english: "",
			russian: "",
			association: "",
			example: "",
			transcription: "",
			image: undefined,
			image_url: "",
			audio: undefined,
			audio_url: ""
		});
  }

  return (
    <>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle sx={{ m: 0, p: 2 }}>
          Add Card
          <IconButton
            aria-label="close"
            onClick={handleClose}
            sx={{
              position: 'absolute',
              right: 8,
              top: 8,
              color: (theme) => theme.palette.grey[500],
            }}
          >
            <CloseIcon/>
          </IconButton>
        </DialogTitle>
        <Box component="form" onSubmit={handleSubmit}>
					<DialogContent dividers>
						<TextField
							label="english"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoComplete="off"
							autoFocus
							onChange={(e) => setFormData(prev => ({...prev, english: e.target.value}))}
							value={formData.english}
						/>
						<TextField
							label="russian"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoComplete="off"
							onChange={(e) => setFormData(prev => ({...prev, russian: e.target.value}))}
							value={formData.russian}
						/>
						<TextField
							label="association"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoComplete="off"
							onChange={(e) => setFormData(prev => ({...prev, association: e.target.value}))}
							value={formData.association}
						/>
						<TextField
							label="example"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoComplete="off"
							onChange={(e) => setFormData(prev => ({...prev, example: e.target.value}))}
							value={formData.example}
						/>
						<TextField
							label="transcription"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoComplete="off"
							onChange={(e) => setFormData(prev => ({...prev, name: e.target.value}))}
							value={formData.transcription}
						/>
          	<Box sx={{display: "flex", justifyContent: "space-between", alignItems: "center"}}>
            	<Box>
              	<Typography>Image</Typography>
								<Box
									onChange={(e) => {
										const target = e.target as HTMLInputElement;
										const file: File = (target.files as FileList)[0];
										setFormData(prev => ({...prev, image: file}));
									}}
									disabled={formData.image_url != undefined && formData.image_url.length > 0} component="input" type="file" accept="image/* "name="image"
								/>
								<Typography>Audio</Typography>
								<Box
									onChange={(e) => {
										const target = e.target as HTMLInputElement;
										const file: File = (target.files as FileList)[0];
										setFormData(prev => ({...prev, audio: file}));
									}}
									disabled={audioChecked || formData.audio_url != undefined && formData.audio_url.length > 0} component="input" type="file" accept="audio/* "name="audio"
								/>
            	</Box>
              <Box>
                <Typography>Load audio from google</Typography>
                <Switch checked={audioChecked} onChange={async (_, checked) => {
									console.log(`/google/audio/${formData.english}`)
                  if (checked) {
                    const res = await api.get(`/google/audio/${formData.english}`);
										const url = res.data?.url;
                    if (res.status === 200) {
                      setAudioChecked(true);
                      setFormData(prev => ({...prev, audio_url: url}));
                      snackbar.enqueueSnackbar("Loaded", {variant: "success"});
                    } else {
                      snackbar.enqueueSnackbar(res.data?.error?.code, {variant: "error"});
                    }
                  } else {
										setAudioChecked(false);
										setFormData(prev => ({...prev, audio_url: ""}));
									}
                }}/>
              </Box>
            </Box>
						<TextField
							label="image url"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoComplete="off"
							onChange={(e) => setFormData(prev => ({...prev, image_url: e.target.value}))}
							value={formData.image_url}
						/>
						<TextField
							label="audio url"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoComplete="off"
							disabled={audioChecked}
							onChange={(e) => setFormData(prev => ({...prev, audio_url: e.target.value}))}
							value={formData.audio_url}
						/>
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose} type="submit">
              Add
            </Button>
          </DialogActions>
        </Box>
      </Dialog>
    {loading && <LoadingBackdrop />}
    </>
  );
}
