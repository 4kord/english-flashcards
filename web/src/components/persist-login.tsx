import React from "react";
import { Outlet } from "react-router-dom";
import { ApiClientMethod, useApi } from "../hooks/use-api";
import { useAuth } from "../hooks/use-auth"
import { LoadingFull } from "./loading-full";

interface IResponse {
  userId: number;
  email: string;
  role: number;
}

export const PersistLogin: React.FC = () => {
  const { auth, setAuth } = useAuth();

  const [fetch, {loading}] = useApi<void, IResponse>({
    method: ApiClientMethod.GET,
    url: "/api/me",
    onSuccess: (data) => {
      console.log(data);
      setAuth(data);
    },
    onFail: (error) => {
      console.log(error)
    },
    defaultLoading: true
  });

  React.useEffect(() => {
    if (!auth) {
      fetch({});
    }
  }, [auth, fetch]);

  return (
    loading
      ? <LoadingFull />
      : <Outlet />
  );
}
