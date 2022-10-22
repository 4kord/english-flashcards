import { Box, Button, TextField, Typography } from "@mui/material";
import { useSnackbar } from "notistack";
import { Controller, SubmitHandler, useForm, useFormState } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { LoadingBackdrop } from "../../../components/loading-backdrop";
import { ApiClientMethod, useApi } from "../../../hooks/use-api";
import { confirmPasswordValidation, emailValidation, passwordValidation } from "../validation";
import { IRequest, IResponse } from "./types";

export const SignupForm = () => {
  const { enqueueSnackbar } = useSnackbar();

  const {handleSubmit, control, watch} = useForm<IRequest>({
    mode: "onBlur"
  });
  const { errors } = useFormState({
    control
  });

  const navigate = useNavigate();

  const [fetch, { loading }] = useApi<IRequest, IResponse>({
    method: ApiClientMethod.POST,
    url: "/api/register",
    onSuccess: (data) => {
      console.log(data);
      enqueueSnackbar("Successfully signed up", {variant: "success"});
      navigate("/signin", {replace: true})
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
    <Box component="form" onSubmit={handleSubmit(submit)} sx={{
      width: 500,
      marginX: 2,
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
    }}>
      <Typography component="h1" variant="h4">Sign Up</Typography>
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
      <Controller
        control={ control }
        name="confirmPassword"
        defaultValue=""
        rules={confirmPasswordValidation(watch("password"))}
        render={({ field }) => (
          <TextField
            type="password"
            label="confirm password"
            size="small"
            margin="normal"
            fullWidth={true}
            onChange={(e) => field.onChange(e)}
            onBlur={() => field.onBlur()}
            value={field.value}
            error={!!errors.confirmPassword?.message}
            helperText={errors.confirmPassword?.message}
          />
        )}
      />
      <Button variant="contained" fullWidth type="submit">Sign Up</Button>
    </Box>
    {loading && <LoadingBackdrop />}
    </>
  )
}
