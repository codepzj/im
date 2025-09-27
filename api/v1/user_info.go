package v1

import (
	"im/internal/config"
	"im/internal/dto/req"
	"im/pkg"

	"github.com/gin-gonic/gin"
)

type UserInfoHandler struct {
	conf *config.Config
}

func NewUserInfoHandler(conf *config.Config) *UserInfoHandler {
	return &UserInfoHandler{
		conf: conf,
	}
}

func (h *UserInfoHandler) GetSmsCode(ctx *gin.Context) {
	var req req.SmsCodeSendReq
	if err := ctx.ShouldBind(&req); err != nil {
		JsonBack(ctx, err.Error(), -2, nil)
		return
	}

	tecentSms := h.conf.TecentSms
	if err := pkg.SendTecentSmsCode(tecentSms.AccessKeyId, tecentSms.AccessKeySecret,
		tecentSms.SmsSdkAppId, tecentSms.SignName, tecentSms.TemplateId,
		"666", req.Phone); err != nil {
		JsonBack(ctx, err.Error(), -1, nil)
		return
	}
	JsonBack(ctx, "发送成功", 0, nil)
}
