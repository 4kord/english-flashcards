import { Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, TextField } from "@mui/material";
import { Box } from "@mui/system";
import CloseIcon from '@mui/icons-material/Close';
import { IDeck } from "../../../../../types/deck";
import { Dispatch, SetStateAction } from "react";
import { Controller, SubmitHandler, useForm, useFormState } from "react-hook-form";
import { IFormData, IRequest } from "./types";
import { ApiClientMethod, useApi } from "../../../../../hooks/use-api";
import { useAuth } from "../../../../../hooks/use-auth";
import { useSnackbar } from "notistack";

interface IProps {
  open: boolean;
  handleClose: () => void;
  setDeckList: Dispatch<SetStateAction<IDeck[]>>
}

export const CreateDeckDialog: React.FC<IProps> = ({open, handleClose, setDeckList}) => {
  const { enqueueSnackbar } = useSnackbar();

  const { auth } = useAuth();

  const [fetch] = useApi<IRequest, IDeck>({
    method: ApiClientMethod.POST,
    url: "/api/decks/premade",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully created deck", {variant: "success"});
      setDeckList((prev) => {
        if (prev) {
          return [...prev, data];
        } else {
          return [data];
        }
      });
    },
    onFail: (error) => {
      enqueueSnackbar(error.errorCode, {variant: "error"});
      console.log(error);
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
      userId: auth.userId,
      name: formData.name,
      isPremade: true,
    }});
  };

  return (
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle sx={{ m: 0, p: 2 }}>
          Create Deck
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
              defaultValue=""
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
