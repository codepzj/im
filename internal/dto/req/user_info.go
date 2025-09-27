package req

type LoginReq struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginWithSmsReq struct {
	Phone     string `json:"phone" binding:"required"`
	PhoneCode string `json:"phoneCode" binding:"required"`
}

type RegisterReq struct {
	NickName  string `json:"nickname" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	PhoneCode string `json:"phoneCode" binding:"required"` // 验证码
}

type SmsCodeSendReq struct {
	Phone string `form:"phone" binding:"required"`
}
