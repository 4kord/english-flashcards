import { Box, Button, Divider, Input, Paper, Typography } from "@mui/material";
import AddIcon from '@mui/icons-material/Add';
import React, { useEffect, useState } from "react";
import { ApiClientMethod, useApi } from "@/hooks/use-api";
import { LoadingSmall } from "@/components/loading-small";
import { useSnackbar } from "notistack";
import { LoadingBackdrop } from "@/components/loading-backdrop";
import { useCalculateHeight } from "@/hooks/use-calculate-height";
import { User } from "@/models/user";
import { AddUserModal } from "./components/add-user-modal";

const UsersBody: React.FC<{users: Array<User>, loading: boolean}> = ({users, loading}) => {
	return (
		<>
		<Box sx={{
			paddingX: 2,
		}}>
			{
				users && users.map((user) => (
					<Paper key={user.id} variant="outlined" sx={{
						padding: 2,
						minHeight: 75,
						display: "flex",
						flexWrap: "nowrap",
						justifyContent: "space-between",
						alignItems: "center",
						marginY: 2
					}}>
						<Typography variant="h5">{`${user.email} (${user.id})`}</Typography>
						<Button color="error" onClick={() => {() => {}}}>Delete</Button>
					</Paper>
				))
			}
		</Box>
		{loading && <LoadingSmall />}
		</>
	);
}

const UsersHeader: React.FC<{handleAddUser: () => void}> = ({handleAddUser}) => {
	return (
		<Box sx={{
			padding: 2,
			display: "flex",
			justifyContent: "space-between",
		}}>
			<Typography component="h3" variant="h5" sx={{ fontWeight: 600 }}>Users</Typography>
			<Button variant="contained" startIcon={<AddIcon />} onClick={handleAddUser}>Add</Button>
		</Box>
	);
}

export const Users: React.FC = () => {
  const snackbar = useSnackbar();
  const calculatedHeight = useCalculateHeight()

  // get users
  const [getUsers, {data: users, setData: setUsers, loading}] = useApi<void, Array<User>>({
    method: ApiClientMethod.GET,
    url: "/users",
    onSuccess: (data) => {
      console.log(data);
      snackbar.enqueueSnackbar("Successfully fetched users", {variant: "success"});
    },
    onFail: (error) => {
      console.log(error);
      snackbar.enqueueSnackbar(error.error.code, {variant: "error"});
    }

  });

  useEffect(() => {
    getUsers();
  }, [getUsers]);

	// Add User Modal
	const [open, setOpen] = useState(false);

  return (
    <>
		<AddUserModal open={open} handleClose={() => setOpen(false)} setUserList={setUsers} />
    {calculatedHeight &&
      <Box sx={{
        overflowY: "scroll",
        height: calculatedHeight
      }}>
				<UsersHeader handleAddUser={() => setOpen(true)} />
        <Divider />
				<UsersBody users={users} loading={loading} />
      </Box>
    }
    {/* {loading && <LoadingBackdrop />} */}
    </>
  );
}
