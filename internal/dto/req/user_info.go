package req

type SmsCodeSendReq struct {
	Phone string `form:"phone" binding:"required"`
}
