import * as React from "react"
import { Auth, AuthContext } from "@/stores/auth-context";

export const useAuth = () => {
    return React.useContext(AuthContext) as {auth: Auth | null, setAuth: React.Dispatch<React.SetStateAction<Auth | null>>};
}
