import axios, { Axios, AxiosError, AxiosRequestHeaders } from "axios";

const URL = import.meta.env.VITE_API_URL;

export const restApiPublic = axios.create({
    baseURL: URL,
    headers: { "Content-Type": "application/json" },
    withCredentials: true,
});

export const restApi = axios.create({
    baseURL: URL,
    headers: { "Content-Type": "application/json" },
    withCredentials: true,
});

restApi.interceptors.request.use((config) => {
    (config.headers as AxiosRequestHeaders)["Authorization"] = `Bearer ${localStorage.getItem("access_token")}`

    return config
})

restApi.interceptors.response.use(
    response => response,
    async error => {
        const originalRequest = error.config;
        if (error.response.status === 401 && !originalRequest._retry) {
          originalRequest._retry = true;
          await refreshAccessToken();
          return restApi(originalRequest);
        }
        return Promise.reject(error);
    }
)

const refreshAccessToken = async () => {
    const res = await restApiPublic.get("/auth/refresh");
    localStorage.setItem("access_token", res.data.access_token);
} 