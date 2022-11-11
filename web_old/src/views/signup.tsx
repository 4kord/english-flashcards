import { Box } from "@mui/material"
import { SignupForm } from "../features/authentication/signup-form"
import { useWindowSize } from "../hooks/use-window-size";

export const SignupView: React.FC = () => {
  const windowSize = useWindowSize();

  return (
    <Box sx={{
      height: windowSize.height,
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
    }}>
      <SignupForm />
    </Box>
  )
}
