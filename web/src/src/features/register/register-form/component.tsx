import { Box, Button, TextField, Typography } from "@mui/material"

export const RegisterForm = () => {
  const handleSubmit = () => {
    console.log("submitted")
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
        // onChange={(e) => field.onChange(e)}
        // onBlur={() => field.onBlur()}
        // value={field.value}
        // error={!!errors.email?.message}
        // helperText={errors.email?.message}
      />
      <TextField
        type="password"
        label="password"
        size="small"
        margin="normal"
        fullWidth
        // onChange={(e) => field.onChange(e)}
        // onBlur={() => field.onBlur()}
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
        // onChange={(e) => field.onChange(e)}
        // onBlur={() => field.onBlur()}
        // value={field.value}
        // error={!!errors.confirmPassword?.message}
        // helperText={errors.confirmPassword?.message}
        />
      <Button variant="contained" type="submit">Sign Up</Button>
    </Box>
    {/* {loading && <LoadingBackdrop />} */}
    </>
  )
}
