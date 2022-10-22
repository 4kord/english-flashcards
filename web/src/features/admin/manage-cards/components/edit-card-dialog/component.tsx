import { Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, TextField, Typography } from "@mui/material";
import { Box } from "@mui/system";
import CloseIcon from '@mui/icons-material/Close';
import React, { Dispatch, SetStateAction } from "react";
import { useSnackbar } from "notistack";
import { Controller, SubmitHandler, useForm, useFormState } from "react-hook-form";
import { IFormData, IRequest } from "./types";
import { ICard } from "../../../../../types/card";
import { ApiClientMethod, useApi } from "../../../../../hooks/use-api";
import { LoadingBackdrop } from "../../../../../components/loading-backdrop";

interface IProps {
  open: boolean;
  handleClose: () => void;
  card: null | ICard;
  setCardList: Dispatch<SetStateAction<ICard[]>>;
}

export const EditCardDialog: React.FC<IProps> = ({open, handleClose, card, setCardList}) => {
  const { enqueueSnackbar } = useSnackbar();

  const {handleSubmit, control, register} = useForm<IFormData>({
    mode: "all",
    shouldUnregister: true,
  });
  const { errors } = useFormState({
    control
  });

  const [editCard, {loading}] = useApi<IRequest, ICard>({
    method: ApiClientMethod.PUT,
    url: "/api/cards",
    headers: {
      "Content-Type": "multipart/form-data"
    },
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully edited card", {variant: "success"})
      setCardList((prev) => prev.map(obj => {
        if (obj.id === data.id) {
          return data
        } else {
          return obj
        }
      }));
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"})
    }
  });

  const submit: SubmitHandler<IFormData> = async (formData) => {
    console.log(formData);
    const requestData: IRequest = {...formData, image: formData.image[0], audio: formData.audio[0]}
    console.log(requestData);
    editCard({
      urlAddition: card?.id.toString(),
      data: requestData
    });
  };

  return (
    <>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle sx={{ m: 0, p: 2 }}>
          Edit Card
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
              name="english"
              defaultValue={card?.english}
              render={({ field }) => (
                <TextField 
                  label="english"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoComplete="off"
                  autoFocus
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.english?.message}
                  helperText={errors.english?.message}
                />
              )}
            />
            <Controller
              control={control}
              name="russian"
              defaultValue={card?.russian || ""}
              render={({ field }) => (
                <TextField 
                  label="russian"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoComplete="off"
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.russian?.message}
                  helperText={errors.russian?.message}
                />
              )}
            />
            <Controller
              control={control}
              name="association"
              defaultValue={card?.association || ""}
              render={({ field }) => (
                <TextField 
                  label="association"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoComplete="off"
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.association?.message}
                  helperText={errors.association?.message}
                />
              )}
            />
            <Controller
              control={control}
              name="example"
              defaultValue={card?.example || ""}
              render={({ field }) => (
                <TextField 
                  label="example"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoComplete="off"
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.example?.message}
                  helperText={errors.example?.message}
                />
              )}
            />
            <Controller
              control={control}
              name="transcription"
              defaultValue={card?.transcription || ""}
              render={({ field }) => (
                <TextField 
                  label="transcription"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoComplete="off"
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.transcription?.message}
                  helperText={errors.transcription?.message}
                />
              )}
            />
            <Box sx={{display: "flex", justifyContent: "space-between", alignItems: "center"}}>
            <Typography>Image</Typography>
            <Box component="input" {...register("image")} type="file" accept="image/* "name="image"/>
            <Typography>Audio</Typography>
            <Box component="input" {...register("audio")} type="file" accept="audio/* "name="audio"/>
            </Box>
            <Controller
              control={control}
              name="imageUrl"
              defaultValue={card?.imageUrl || ""}
              render={({ field }) => (
                <TextField 
                  label="image url"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoComplete="off"
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.imageUrl?.message}
                  helperText={errors.imageUrl?.message}
                />
              )}
            />
            <Controller
              control={control}
              name="audioUrl"
              defaultValue={card?.audioUrl || ""}
              render={({ field }) => (
                <TextField 
                  label="audio url"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoComplete="off"
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.audioUrl?.message}
                  helperText={errors.audioUrl?.message}
                />
              )}
            />
          </DialogContent>
          <DialogActions>
            <Button type="submit" onClick={handleClose}>
              Edit
            </Button>
          </DialogActions>
        </Box>
      </Dialog>
    {loading && <LoadingBackdrop />}
    </>
  );
}
