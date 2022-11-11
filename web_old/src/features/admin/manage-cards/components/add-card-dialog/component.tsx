import { Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, Switch, TextField, Typography } from "@mui/material";
import { Box } from "@mui/system";
import CloseIcon from '@mui/icons-material/Close';
import React, { Dispatch, SetStateAction } from "react";
import { Controller, SubmitHandler, useForm, useFormState } from "react-hook-form";
import { IFormData, IRequest } from "./types";
import { ICard } from "../../../../../types/card";
import { ApiClientMethod, useApi } from "../../../../../hooks/use-api";
import { useSnackbar } from "notistack";
import { LoadingBackdrop } from "../../../../../components/loading-backdrop";
import { axiosInstance } from "../../../../../lib/axios";

interface IProps {
  open: boolean;
  handleClose: () => void;
  setCardList: Dispatch<SetStateAction<ICard[]>>;
  deckId: string;
}

export const AddCardDialog: React.FC<IProps> = ({open, handleClose, setCardList, deckId}) => {
  const { enqueueSnackbar } = useSnackbar();

  const {handleSubmit, control, register, reset, getValues, setValue} = useForm<IFormData>({
    mode: "all",
    shouldUnregister: true,
  });
  const { errors } = useFormState({
    control
  });

  const [fetch, {loading}] = useApi<IRequest, ICard>({
    method: ApiClientMethod.POST,
    url: "/api/decks",
    headers: {
      "Content-Type": "multipart/form-data"
    },
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully added card", {variant: "success"})
      setCardList((prev) => {
        if (prev) {
          return [...prev, data];
        } else {
          return [data];
        }
      })
      reset()
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
    await fetch({
      data: requestData,
      urlAddition: deckId
    });
  };

  const [audioChecked, setAudioChecked] = React.useState<boolean>(false);

  return (
    <>
      <Dialog open={open} onClose={() => {
        handleClose();
        setAudioChecked(false);
      }}>
        <DialogTitle sx={{ m: 0, p: 2 }}>
          Add Card
          <IconButton
            aria-label="close"
            onClick={() => {
              handleClose();
              setAudioChecked(false);
            }}
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
              defaultValue=""
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
              defaultValue=""
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
              defaultValue=""
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
              defaultValue=""
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
              defaultValue=""
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
            <Box>
              <Typography>Image</Typography>
              <Box component="input" {...register("image")} type="file" accept="image/* "name="image"/>
              <Typography>Audio</Typography>
              <Box disabled={audioChecked} component="input" {...register("audio")} type="file" accept="audio/* "name="audio"/>
            </Box>
              <Box>
                <Typography>Load audio from google</Typography>
                <Switch checked={audioChecked} onChange={async (_, checked) => {
                  if (checked) {
                    const res = await axiosInstance.get(`/api/google/audio/${getValues("english")}`);
                    if (res.data?.status === 200) {
                      setAudioChecked(true);
                      setValue("audioUrl", res.data?.url)
                      enqueueSnackbar("Loaded", {variant: "success"})
                      return
                    } else {
                      enqueueSnackbar("Audio not found", {variant: "error"})
                    }
                  }
                  setAudioChecked(false);
                  setValue("audioUrl", "")
                }}/>
              </Box>
            </Box>
            <Controller
              control={control}
              name="imageUrl"
              defaultValue=""
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
              defaultValue=""
              render={({ field }) => (
                <TextField 
                  label="audio url"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoComplete="off"
                  disabled={audioChecked}
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
            <Button type="submit" onClick={() => setAudioChecked(false)}>
              Add
            </Button>
          </DialogActions>
        </Box>
      </Dialog>
    {loading && <LoadingBackdrop />}
    </>
  );
}
