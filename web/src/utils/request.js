import axios from "axios";

// 创建 axios 实例
const request = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
request.interceptors.response.use(
  (response) => {
    const { data } = response;
    if (data.code === 200) {
      return data;
    } else {
      return Promise.reject(new Error(data.message || "请求失败"));
    }
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default request;
