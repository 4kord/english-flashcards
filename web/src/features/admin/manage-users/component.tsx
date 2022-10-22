import { Button, Divider, Input, Paper, Typography } from "@mui/material";
import { Box } from "@mui/system";
import AddIcon from '@mui/icons-material/Add'; 
import React, { useEffect } from "react";
import { ApiClientMethod, useApi } from "../../../hooks/use-api";
import { LoadingSmall } from "../../../components/loading-small";
import { AddUserDialog } from "./components/add-user-dialog";
import { IUserResponsePayload } from "../../../types/user";
import { useSnackbar } from "notistack";
import { LoadingBackdrop } from "../../../components/loading-backdrop";
import { useCalculateHeight } from "../../../hooks/use-calculate-height";

export const ManageUsers: React.FC = () => {
  const { enqueueSnackbar } = useSnackbar();

  // get users
  const [getUsers, {data, setData, loading}] = useApi<void, Array<IUserResponsePayload>>({
    method: ApiClientMethod.GET,
    url: "/api/users",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully fetched users", {variant: "success"});
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }

  });

  useEffect(() => {
    getUsers({});
  }, [getUsers]);

  // search
  const [search, setSearch] = React.useState<string>("");

  // open and close dialog
  const [dialogOpen, setDialogOpen] = React.useState(false);

  const handleClickOpenDialog = () => {
    setDialogOpen(true);
  };

  const handleCloseDialog = () => {
    setDialogOpen(false);
  };

  // delete user
  const [deleteUser, { loading: deleteUserLoading }] = useApi<void, void>({
    method: ApiClientMethod.DELETE,
    url: "/api/users",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully deleted user", {variant: "success"})
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"})
    }
  });
  const deleteHandler = async (userId: number) => {
    console.log(userId);
    await deleteUser({urlAddition: userId.toString()});
    setData((prev) => prev.filter(user => user.userId !== userId));
  }

  const calculatedHeight = useCalculateHeight()

  return (
    <>
    <AddUserDialog open={dialogOpen} handleClose={handleCloseDialog} setUserList={setData} />
    {calculatedHeight &&
      <Box sx={{
        overflowY: "scroll",
        height: calculatedHeight
      }}>
        <Box sx={{
          padding: 2,
          display: "flex",
          justifyContent: "space-between",
        }}>
          <Typography component="h3" variant="h5" sx={{ fontWeight: 600 }}>Users</Typography>
          <Button variant="contained" startIcon={<AddIcon />} onClick={handleClickOpenDialog}>Add</Button>
        </Box>
        <Divider /> <Box>
          <Input placeholder="Search" fullWidth onChange={(e: React.ChangeEvent<HTMLInputElement>) => setSearch(e.target.value)} sx={{
            padding: 2
          }}/>
        </Box>
        <Box sx={{
          padding: 2,
        }}>
          {
            data && data.filter((user) => user.email.includes(search) || user.userId.toString() === search).map((user) => (
              <Paper key={user.userId} variant="outlined" sx={{
                padding: 2,
                minHeight: 75,
                display: "flex",
                flexWrap: "nowrap",
                justifyContent: "space-between",
                alignItems: "center",
                marginY: 2
              }}>
                <Typography variant="h5">{`${user.email} (${user.userId})`}</Typography>
                <Button color="error" onClick={() => {deleteHandler(user.userId)}}>Delete</Button>
              </Paper>
            ))
          }
          {loading && <LoadingSmall />}
        </Box>
      </Box>
    }
    {deleteUserLoading && <LoadingBackdrop />}
    </>
  );
}
