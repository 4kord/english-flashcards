import { Box, Button, TextField, Typography } from "@mui/material";
import { useSnackbar } from "notistack";
import { Controller, SubmitHandler, useForm, useFormState } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { LoadingBackdrop } from "../../../components/loading-backdrop";
import { ApiClientMethod, useApi } from "../../../hooks/use-api";
import { useAuth } from "../../../hooks/use-auth";
import { IUserResponsePayload } from "../../../types/user";
import { emailValidation, passwordValidation } from "../validation";
import { IRequest } from "./types";

export const SigninForm = () => {
  const { enqueueSnackbar } = useSnackbar();

  const {handleSubmit, control} = useForm<IRequest>({
    mode: "onBlur"
  });
  const { errors } = useFormState({
    control
  });

  const { setAuth } = useAuth();
  const navigate = useNavigate();

  const [fetch, {loading}] = useApi<IRequest, IUserResponsePayload>({
    method: ApiClientMethod.POST,
    url: "/api/login",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully signed in", {variant: "success"});
      setAuth(data);
      navigate("/", {replace: true})
    },
    onFail: (error) => {
      console.log(error);
      enqueueSnackbar(error.errorCode, {variant: "error"});
    }
  });

  const submit: SubmitHandler<IRequest> = async (formData) => {
    await fetch({data: formData})
  };

  return (
    <>
    <Box component="form" onSubmit={handleSubmit(submit)} sx={{
      width: 500,
      marginX: 2,
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
    }}>
      <Typography component="h1" variant="h4" >Sign In</Typography>
      <Controller
        control={ control }
        name="email"
        defaultValue=""
        rules={emailValidation}
        render={({ field }) => (
          <TextField
            label="email"
            size="small"
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
        rules={passwordValidation}
        render={({ field }) => (
          <TextField
            type="password"
            label="password"
            size="small"
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
      <Button variant="contained" type="submit" fullWidth>Sign In</Button>
    </Box>
    {loading && <LoadingBackdrop />}
    </>
  )
}
