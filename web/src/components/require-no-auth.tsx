import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "@/hooks/use-auth";

export const RequireNoAuth: React.FC = () => {
    const { auth } = useAuth();

    return (
        !auth
            ? <Outlet />
            : <Navigate to={"/"} />
    );
}
