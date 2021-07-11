import axios from "axios";
import Cookies from "js-cookie";

const token = Cookies.get("token");
axios.defaults.baseURL = `${import.meta.env.VITE_APP_API}`;
axios.defaults.headers.common['Authorization'] = "Bearer " + token;


const _axios = axios.create({});
_axios.interceptors.request.use(
    function (config) {
        // Do something before request is sent
        return config;
    },
    function (error) {
        // Do something with request error
        return Promise.reject(error);
    }
);

// Add a response interceptor
_axios.interceptors.response.use(
    function (response) {
        // Do something with response data
        return response;
    },
    function (error) {
        // Do something with response error
        if (null !== error) {
            alert(error.response.data.error || "未知错误")

            if (error.response.status === 401) {
                // window.location.href = '/login.html'
                if (error.response.data.data) {
                    window.location.href = error.response.data.data
                }
                return
            }
            if (error.response.status === 403) {
                // window.location.href = '/403.html'
            }
        }
        return Promise.reject(error);
    }
);


export default _axios;