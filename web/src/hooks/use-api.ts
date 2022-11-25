import * as React from "react";
import axios, { AxiosError, AxiosRequestHeaders } from "axios";
import { api } from "@/lib/axios";

export enum ApiClientMethod {
  GET = "get",
  POST = "post",
  PUT = "put",
  PATCH = "patch",
  DELETE = "delete"
}

interface IErrorResponse {
  error: {
    kind: string,
    code: string,
    message: string,
  }
}

export const useApi = <D, T>({method, url, headers, onSuccess, onFail, defaultLoading = false}: {method: ApiClientMethod, url: string, headers?: AxiosRequestHeaders, onSuccess?: (data: T) => void, onFail?: (error: IErrorResponse) => void, defaultLoading?: boolean}): [(data?: D) => Promise<void>, {data: T, setData: React.Dispatch<React.SetStateAction<T>>, error: IErrorResponse, loading: boolean}] => {
  const [controller, setController] = React.useState<AbortController | null>(null);
  const [returnData, setReturnData] = React.useState<T | null>(null);
  const [error, setError] = React.useState<IErrorResponse | null>(null);
  const [loading, setLoading] = React.useState<boolean>(defaultLoading);

  const fetch = React.useCallback(async (data?: D) => {
    try {
      setLoading(true);
      const ctrl = new AbortController();
      setController(ctrl);

      const response = await api.request({
        method: method,
        headers: headers,
        url: url,
        data: data,
        signal: ctrl.signal
      });

      setReturnData(response.data);
      onSuccess && onSuccess(response.data);
    } catch (error) {
      const err = error as AxiosError;
      setError((err.response?.data) as IErrorResponse);
      onFail && onFail(err.response?.data as IErrorResponse);
      throw(error)
    } finally {
      setLoading(false);
    }
  // eslint-disable-next-line
  }, [method, url])

  React.useEffect(() => {
    return () => {
      controller && controller.abort();
    }
  }, [controller]);

  return [
    fetch,
    {
      data: returnData!,
      setData: setReturnData as React.Dispatch<React.SetStateAction<T>>,
      error: error!,
      loading: loading,
    }
  ]
}
