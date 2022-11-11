import { Box } from "@mui/material"
import { SigninForm } from "../features/authentication/signin-form"
import { useWindowSize } from "../hooks/use-window-size"

export const SigninView: React.FC = () => {
  const windowSize = useWindowSize();

  return (
    <Box sx={{
      height: windowSize.height,
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
    }}>
      <SigninForm />
    </Box>
  )
}
