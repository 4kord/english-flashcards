import { Button, Divider, Paper, Typography } from "@mui/material";
import { Box } from "@mui/system";
import { useSnackbar } from "notistack";
import { useEffect } from "react";
import { ApiClientMethod, useApi } from "../hooks/use-api";
import { useCalculateHeight } from "../hooks/use-calculate-height";
import AddIcon from '@mui/icons-material/Add'; 
import { LoadingSmall } from "../components/loading-small";
import { useAuth } from "../auth/use-auth";

interface IResponseEntity {
    id: string;
    email: string;
    admin: string;
    created_at: string;
}

export const AdminUsersPage: React.FC = () => {
    const { enqueueSnackbar } = useSnackbar();
    const calculatedHeight = useCalculateHeight();
    const { auth } = useAuth();

    // get users
    const [getUsers, {data, setData, loading}] = useApi<{}, Array<IResponseEntity>>({
      method: ApiClientMethod.GET,
      url: "/users",
      onSuccess: (data) => {
        console.log(data);
        enqueueSnackbar("Successfully fetched users", {variant: "success"});
      },
      onFail: (error) => {
        console.log(error);
        enqueueSnackbar(error.error.code, {variant: "error"});
      },
    });
  
    useEffect(() => {
      getUsers({});
    }, [getUsers]);
  
    return (
      <>
            { data &&
              data.map((user) => (
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
      </>
    );
  }  