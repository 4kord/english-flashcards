import { Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, TextField } from "@mui/material";
import { Box } from "@mui/system";
import CloseIcon from '@mui/icons-material/Close';
import React, { Dispatch, FormEvent, SetStateAction, useState } from "react";
import { ApiClientMethod, useApi } from "@/hooks/use-api";
import { LoadingBackdrop } from "@/components/loading-backdrop";
import { useSnackbar } from "notistack";
import { AddDeck, Deck } from "@/models/deck";
import { useAuth } from "@/hooks/use-auth";

interface Props {
  open: boolean;
  handleClose: () => void;
  setDeckList: Dispatch<SetStateAction<Deck[]>>;
}

export const AddDeckModel: React.FC<Props> = ({ open, handleClose, setDeckList }) => {
	const { auth } = useAuth();
  const snackbar = useSnackbar();

	const [formData, setFormData] = useState<AddDeck>({
		name: ""
	})

  const [addDeck, { loading }] = useApi<AddDeck, Deck>({
    method: ApiClientMethod.POST,
    url: `/users/${auth?.user_id}/decks/premade`,
    onSuccess: (data) => {
      console.log(data);
      snackbar.enqueueSnackbar("Successfully added deck", {variant: "success"});
      setDeckList(prev => {
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

    await addDeck({
			data: formData
		})

    setFormData({name: ""});
  }

  return (
    <>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle sx={{ m: 0, p: 2 }}>
          Add Deck
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
							label="name"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoFocus
							onChange={(e) => setFormData(prev => ({...prev, name: e.target.value}))}
							value={formData.name}
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
