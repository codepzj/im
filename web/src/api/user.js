import request from "@/utils/request";

// 发送验证码
export const sendSmsCode = (phone) => {
  return request.get("/send/smsCode", { params: { phone } });
};

// 用户登录
export const login = (data) => {
  return request.post("/login", data);
};

// 验证码登录
export const loginWithSms = (data) => {
  return request.post("/loginWithSms", data);
};

// 用户注册
export const register = (data) => {
  return request.post("/register", data);
};
