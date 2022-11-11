import { Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, TextField } from "@mui/material";
import { Box } from "@mui/system";
import CloseIcon from '@mui/icons-material/Close';
import { IDeck } from "../../../../../types/deck";
import { Dispatch, SetStateAction } from "react";
import { Controller, SubmitHandler, useForm, useFormState } from "react-hook-form";
import { IFormData, IRequest } from "./types";
import { ApiClientMethod, useApi } from "../../../../../hooks/use-api";

interface IProps {
  open: boolean;
  handleClose: () => void;
  deck: IDeck;
  setDeckList: Dispatch<SetStateAction<IDeck[]>>;
}

export const EditDeckDialog: React.FC<IProps> = ({open, handleClose, deck, setDeckList}) => {
  const [fetch] = useApi<IRequest, IDeck>({
    method: ApiClientMethod.PUT,
    url: "/api/decks",
    onSuccess: (data) => {
      console.log(data);
      setDeckList((prev) => prev.map(obj => {
        if (obj.id === data.id) {
          return data
        } else {
          return obj
        }
      }));
    },
    onFail: (error) => {
      console.error(error);
    }
  });

  const {handleSubmit, control} = useForm<IFormData>({
    mode: "onBlur",
    shouldUnregister: true,
  });
  const { errors } = useFormState({
    control
  });

  const submit: SubmitHandler<IFormData> = async (formData) => {
    await fetch({data: {
      name: formData.name,
    },
      urlAddition: deck.id.toString()
    });
  };

  return (
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle sx={{ m: 0, p: 2 }}>
          Edit Deck
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
        <Box component="form" onSubmit={handleSubmit(submit)}>
          <DialogContent dividers>
            <Controller
              control={control}
              name="name"
              defaultValue={deck?.name}
              render={({ field }) => (
                <TextField 
                  label="name"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.name?.message}
                  helperText={errors.name?.message}
                />
              )}
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose} type="submit" autoFocus>
              Add
            </Button>
          </DialogActions>
        </Box>
      </Dialog>
  );
}
