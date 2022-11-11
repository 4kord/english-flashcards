import { Box } from "@mui/material";
import { Link } from "react-router-dom";
import { useAuth } from "../auth/use-auth"

export const HomePage: React.FC = () => {
    const { auth } = useAuth();

    return (
        <Box>
            <div>Index</div>
            {auth?.email}
            <Link to="/login">Login</Link>
            <Link to="/register">Register</Link>
        </Box>
    )
}