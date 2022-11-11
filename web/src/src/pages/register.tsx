import { Box } from "@mui/system"
import { useWindowSize } from "../hooks/use-window-size";
import { RegisterForm } from "../features/register/register-form"

export const RegisterPage: React.FC = () => {
    const windowSize = useWindowSize();

    return (
        <Box sx={{
            width: "100%",
            height: windowSize.height,
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            paddingX: 3,
        }}>
            <RegisterForm />
        </Box>
    )
}