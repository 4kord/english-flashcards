import { Button, Dialog, DialogActions, DialogContent, DialogTitle, IconButton, MenuItem, TextField } from "@mui/material";
import { Box } from "@mui/system";
import CloseIcon from '@mui/icons-material/Close';
import React, { Dispatch, SetStateAction } from "react";
import { IRequest } from "./types";
import { Controller, SubmitHandler, useForm, useFormState } from "react-hook-form";
import { IUserResponsePayload } from "../../../../../types/user";
import { ApiClientMethod, useApi } from "../../../../../hooks/use-api";
import { LoadingBackdrop } from "../../../../../components/loading-backdrop";
import { useSnackbar } from "notistack";

interface IProps {
  open: boolean;
  handleClose: () => void;
  setUserList: Dispatch<SetStateAction<IUserResponsePayload[]>>
}

export const AddUserDialog: React.FC<IProps> = ({ open, handleClose, setUserList }) => {
  const { enqueueSnackbar } = useSnackbar();

  const roles = [
    {
      value: 1,
      label: "Admin"
    },
    {
      value: 2,
      label: "User"
    }
  ]

  const {handleSubmit, control} = useForm<IRequest>({
    mode: "onBlur",
    shouldUnregister: true,
  });
  const { errors } = useFormState({
    control
  });

  const [fetch, { loading }] = useApi<IRequest, IUserResponsePayload>({
    method: ApiClientMethod.POST,
    url: "/api/users",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully added user", {variant: "success"});
      setUserList(prev => {
        if (prev) {
          return [...prev, data];
        } else {
          return [data];
        }
      });
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  const submit: SubmitHandler<IRequest> = async (formData) => {
    await fetch({data: formData});
  };

  return (
    <>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle sx={{ m: 0, p: 2 }}>
          Add User
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
              control={ control }
              name="email"
              defaultValue=""
              render={({ field }) => (
                <TextField 
                  label="email"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  autoFocus
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.email?.message}
                  helperText={errors.email?.message}
                />
              )}
            />
            <Controller
              control={ control }
              name="password"
              defaultValue=""
              render={({ field }) => (
                <TextField 
                  label="password"
                  type="password"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.password?.message}
                  helperText={errors.password?.message}
                />
              )}
            />
            <Controller
              control={ control }
              name="role"
              defaultValue={2}
              render={({ field }) => (
                <TextField 
                  select
                  label="role"
                  size="medium"
                  margin="normal"
                  fullWidth={true}
                  defaultValue=""
                  onChange={(e) => field.onChange(e)}
                  onBlur={() => field.onBlur()}
                  value={field.value}
                  error={!!errors.role?.message}
                  helperText={errors.role?.message}
                >
                  {roles.map((option) => (
                    <MenuItem key={option.value} value={option.value}>
                      {option.label}
                    </MenuItem>
                  ))}
                </TextField>
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
    {loading && <LoadingBackdrop />}
    </>
  );
}
