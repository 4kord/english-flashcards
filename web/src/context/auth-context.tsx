import React from "react";

export interface Auth {
  email: string
  userId: number
  role: number
}

export const AuthContext = React.createContext<{auth: Auth | null, setAuth: React.Dispatch<React.SetStateAction<Auth | null>> | null}>({auth: null, setAuth: null});

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
