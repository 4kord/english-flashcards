import React from "react";
import { Outlet } from "react-router-dom";
import { ApiClientMethod, useApi } from "../hooks/use-api";
import { useAuth } from "./use-auth";
import { LoadingFull } from "../components/loading-full";
import { AxiosRequestConfig } from "axios";

interface IResponse {
  userID: number;
	email: string;
	admin: boolean;
}

export const User: React.FC = () => {
  const { auth, setAuth } = useAuth();

  const [fetch, {loading, data}] = useApi<{}, IResponse>({
    method: ApiClientMethod.GET,
    url: "/auth/user",
    onSuccess: (data) => {
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
    loading && !data
      ? <LoadingFull />
      : <Outlet />
  );
}
