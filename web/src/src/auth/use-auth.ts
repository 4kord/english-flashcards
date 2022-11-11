import React from "react"
import { Auth, AuthContext } from "./store";

export const useAuth = () => {
    return React.useContext(AuthContext) as {auth: Auth | null, setAuth: React.Dispatch<React.SetStateAction<Auth | null>>};
}
