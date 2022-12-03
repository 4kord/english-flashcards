import { Box, Button, TextField, Typography } from "@mui/material"
import { useSnackbar } from "notistack";
import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { LoadingBackdrop } from "@/components/loading-backdrop";
import { ApiClientMethod, useApi } from "@/hooks/use-api";

interface RegisterRequest {
  email: string;
  password: string;
}

export const RegisterForm = () => {
  const snackbar = useSnackbar();
  const navigate = useNavigate();
  const [formData, setFormData] = useState<RegisterRequest>({
    email: "",
    password: ""
  });

  const [register, { loading }] = useApi<RegisterRequest, void>({
    method: ApiClientMethod.POST,
    url: "/auth/register",
    onSuccess: (data) => {
      console.log(data);
      snackbar.enqueueSnackbar("Successfully signed up", {variant: "success"});
      navigate("/login", {replace: true})
    },
    onFail: (error) => {
      console.log(error);
      snackbar.enqueueSnackbar(error?.error?.code, {variant: "error"});
    }
  });

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()

    await register({
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
      <Typography component="h1" variant="h4" >Sign Up</Typography>
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
        // value={field.value}
        // error={!!errors.password?.message}
        // helperText={errors.password?.message}
      />
      <TextField
        type="password"
        label="confirm password"
        size="small"
        margin="normal"
        fullWidth
        // onChange={e => setFormData(prev => ({...prev, email: e.target.value}))}
        // value={formData.email}
        // error={!!errors.confirmPassword?.message}
        // helperText={errors.confirmPassword?.message}
        />
      <Button variant="contained" type="submit">Sign Up</Button>
    </Box>
    {loading && <LoadingBackdrop />}
    </>
  )
}
