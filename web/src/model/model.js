import axios from "../plugins/axios";

export const apiShorten = (form) => axios.post(`/api/shorten`, form);
export const apiListUrl = () => axios.get(`/api/urls`)
export const apiConfig = () => axios.get(`/api/configs`)
export const apiLogin = (token) => axios.post(`/auth/login`, {token: token})
export const apiAnalytics = (id, query) => axios.get(`/api/urls/${id}/analytics`, {params: query})