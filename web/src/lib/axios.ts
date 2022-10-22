import axios from "axios";

const URL = process.env.REACT_APP_API

export const axiosInstance = axios.create({
    baseURL: URL,
    headers: { "Content-Type": "application/json" },
    withCredentials: true
})
