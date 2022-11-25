import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "@/hooks/use-auth";

interface Props {
  adminOnly: boolean
}

export const RequireAuth: React.FC<Props> = ({adminOnly}) => {
    const { auth } = useAuth();

    if (!auth) {
        return <Navigate to="/login" /> 
    }

    if (auth.admin === false && adminOnly === true) {
        return <h1>Forbidden</h1> 
    }

    return <Outlet />
}