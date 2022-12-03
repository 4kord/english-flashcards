import { Box, Button, TextField, Typography } from "@mui/material";
import { useSnackbar } from "notistack";
import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ApiClientMethod, useApi } from "@/hooks/use-api";
import { useAuth } from "@/hooks/use-auth";
import { LoadingBackdrop } from "@/components/loading-backdrop";

interface LoginRequest {
  email: string;
  password: string;
}

interface LoginResponse {
  user_id: number;
	email: string;
	admin: boolean;
	access_token: string;
	refresh_token: string;
}

export const LoginForm = () => {
  const snackbar = useSnackbar();
  const navigate = useNavigate();
  const { setAuth } = useAuth();

  const [formData, setFormData] = useState<LoginRequest>({
    email: "",
    password: ""
  });

  const [login, { loading }] = useApi<LoginRequest, LoginResponse>({
    method: ApiClientMethod.POST,
    url: "/auth/login",
    onSuccess: (data) => {
      console.log(data);
      setAuth({
				user_id: data.user_id,
				email: data.email,
				admin: data.admin,
			});
      localStorage.setItem("access_token", data.access_token)
      snackbar.enqueueSnackbar("Successfully signed in", {variant: "success"});
      navigate("/", {replace: true})
    },
    onFail: (error) => {
      console.log(error);
      snackbar.enqueueSnackbar(error?.error?.code, {variant: "error"});
    }
  });

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    console.log(formData);

    await login({
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
    {loading && <LoadingBackdrop />}
    </>
  )
}
