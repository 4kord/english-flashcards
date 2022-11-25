import React from "react";

export interface Auth {
  user_id: number;
	email: string;
	admin: boolean;
}

export const AuthContext = React.createContext<{auth: Auth | null, setAuth: React.Dispatch<React.SetStateAction<Auth | null>> | null}>({auth: null, setAuth: null});
