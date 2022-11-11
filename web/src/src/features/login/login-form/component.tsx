import { Box, Button, TextField, Typography } from "@mui/material";
import { useSnackbar } from "notistack";
import { FormEvent, Reducer, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ApiClientMethod, useApi } from "../../../hooks/use-api";
import { useAuth } from "../../../auth/use-auth";

interface IRequest {
  email: string;
  password: string;
}

interface IResponse {
  userID: number;
	email: string;
	admin: boolean;
	access_token: string;
	refresh_token: string;
}

export const LoginForm = () => {
  const { setAuth } = useAuth();

  const { enqueueSnackbar } = useSnackbar();

  const navigate = useNavigate();

  const [formData, setFormData] = useState<IRequest>({
    email: "",
    password: ""
  });

  const [fetch] = useApi<IRequest, IResponse>({
    method: ApiClientMethod.POST,
    url: "/auth/login",
    onSuccess: (data) => {
      console.log(data);
      setAuth(data)
      localStorage.setItem("access_token", data.access_token)
      enqueueSnackbar("Successfully signed in", {variant: "success"});
      navigate("/", {replace: true})
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error?.error?.code, {variant: "error"});
    }
  });

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    console.log(formData);

    fetch({
      data: formData
    })

    setFormData({email: "", password: ""});
  }

  return (
    <>
    <Box component="form" onSubmit={handleSubmit} sx={{
      width: 500,
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
    }}>
      <Typography component="h1" variant="h4" >Sign In</Typography>
      <TextField
        label="email"
        size="small"
        margin="normal"
        fullWidth
        autoFocus
        onChange={e => setFormData(prev => ({...prev, email: e.target.value}))}
        value={formData.email}
        // error={!!errors.email?.message}
        // helperText={errors.email?.message}
      />
      <TextField
        type="password"
        label="password"
        size="small"
        margin="normal"
        fullWidth
        onChange={e => setFormData(prev => ({...prev, password: e.target.value}))}
        value={formData.password}
        // error={!!errors.password?.message}
        // helperText={errors.password?.message}
      />
      <Button variant="contained" type="submit">Sign In</Button>
    </Box>
    {/* {loading && <LoadingBackdrop />} */}
    </>
  )
}
