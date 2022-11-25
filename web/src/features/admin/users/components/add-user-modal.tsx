import { Button, Checkbox, Dialog, DialogActions, DialogContent, DialogTitle, FormControlLabel, IconButton, MenuItem, TextField } from "@mui/material";
import { Box } from "@mui/system";
import CloseIcon from '@mui/icons-material/Close';
import React, { Dispatch, FormEvent, SetStateAction, useState } from "react";
import { ApiClientMethod, useApi } from "@/hooks/use-api";
import { LoadingBackdrop } from "@/components/loading-backdrop";
import { useSnackbar } from "notistack";
import { AddUser, User } from "@/models/user";

interface Props {
  open: boolean;
  handleClose: () => void;
  setUserList: Dispatch<SetStateAction<User[]>>;
}

export const AddUserModal: React.FC<Props> = ({ open, handleClose, setUserList }) => {
  const snackbar = useSnackbar();

	const [formData, setFormData] = useState<AddUser>({
		email: "",
		password: "",
		is_admin: false
	})

  const [addUser, { loading }] = useApi<AddUser, User>({
    method: ApiClientMethod.POST,
    url: "/api/users",
    onSuccess: (data) => {
      console.log(data);
      snackbar.enqueueSnackbar("Successfully added user", {variant: "success"});
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
      snackbar.enqueueSnackbar(error.error.code, {variant: "error"});
    }
  });

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    console.log(formData);

    // await addUser(formData)

    setFormData({email: "", password: "", is_admin: false});
  }

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
        <Box component="form" onSubmit={handleSubmit}>
          <DialogContent dividers>
						<TextField
							label="email"
							size="medium"
							margin="normal"
							fullWidth={true}
							autoFocus
							onChange={(e) => setFormData(prev => ({...prev, email: e.target.value}))}
							value={formData.email}
						/>
						<TextField
							label="password"
							type="password"
							size="medium"
							margin="normal"
							fullWidth={true}
							onChange={(e) => setFormData(prev => ({...prev, password: e.target.value}))}
							value={formData.password}
						/>
						<FormControlLabel control={<Checkbox onChange={(e) => setFormData(prev => ({...prev, is_admin: e.target.checked}))} checked={formData.is_admin} />} label="Admin" />
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
