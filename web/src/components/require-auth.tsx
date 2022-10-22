import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../hooks/use-auth";
import { ForbiddenView } from "../views/forbidden";

interface Props {
  allowedRoles: Array<number>
}

export const RequireAuth: React.FC<Props> = ({allowedRoles}) => {
    const { auth } = useAuth();

    return (
        allowedRoles?.includes(auth?.role)
            ? <Outlet />
            : auth
                ? <ForbiddenView />
                : <Navigate to="/signin" />
    );
}
