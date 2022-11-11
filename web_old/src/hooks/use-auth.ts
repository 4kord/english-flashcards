import React from "react"
import { Auth, AuthContext } from "../context/auth-context";

export const useAuth = () => {
    return React.useContext(AuthContext) as {auth: Auth, setAuth: React.Dispatch<React.SetStateAction<Auth | null>>};
}
