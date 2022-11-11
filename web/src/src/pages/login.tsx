import { Box } from "@mui/material"
import { useAuth } from "../auth/use-auth";
import { useWindowSize } from "../hooks/use-window-size";
import { LoginForm } from "../features/login/login-form"

export const LoginPage: React.FC = () => {
    const windowSize = useWindowSize();

    const { auth } = useAuth();

    return (
        <Box sx={{
            width: "100%",
            height: windowSize.height,
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            paddingX: 3,
        }}>
            <LoginForm />
            {auth?.email}
        </Box>
    )
}