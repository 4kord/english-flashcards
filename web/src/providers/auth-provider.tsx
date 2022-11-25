import * as React from "react";
import { Auth, AuthContext } from "@/stores/auth-context";

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