import { RegisterForm } from "@/features/auth";
import { useWindowSize } from "@/hooks/use-window-size";
import { Box } from "@mui/material";

export const RegisterPage = () => {
  const windowSize = useWindowSize();

  return (
    <Box sx={{
      height: windowSize.height,
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
    }}>
      <RegisterForm />
    </Box>
  )
}
