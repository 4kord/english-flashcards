import React, { useEffect } from "react";
import { Auth, AuthContext } from "./store";
import { restApi } from "../apis/calls";
import { AxiosRequestConfig } from "axios";

interface Props {
  children: React.ReactNode
}
  
export const AuthProvider: React.FC<Props> = ({ children }) => {
  const [auth, setAuth] = React.useState<Auth | null>(null);

  return (
    <AuthContext.Provider value={{ auth, setAuth }}>
      {children} 
    </AuthContext.Provider>
  );
}