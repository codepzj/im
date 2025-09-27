package v1

import (
	"im/internal/dto/req"
	"im/internal/service"

	"github.com/gin-gonic/gin"
)

type UserInfoHandler struct {
	userInfoService *service.UserInfoService
}

func NewUserInfoHandler(userInfoService *service.UserInfoService) *UserInfoHandler {
	return &UserInfoHandler{
		userInfoService: userInfoService,
	}
}

// 登录
func (h *UserInfoHandler) Login(ctx *gin.Context) {
	var req req.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		JsonBack(ctx, err.Error(), -2, nil)
		return
	}
	// 校验用户是否存在
	err := h.userInfoService.Login(req)
	if err != nil {
		JsonBack(ctx, err.Error(), -1, nil)
		return
	}
	JsonBack(ctx, "登录成功", 0, nil)
}

// 验证码登录
func (h *UserInfoHandler) LoginWithSmsCode(ctx *gin.Context) {
	var req req.LoginWithSmsReq
	if err := ctx.ShouldBind(&req); err != nil {
		JsonBack(ctx, err.Error(), -2, nil)
		return
	}
	err := h.userInfoService.LoginWithSms(req)
	if err != nil {
		JsonBack(ctx, err.Error(), -1, nil)
		return
	}
	JsonBack(ctx, "登录成功", 0, nil)
}

// 发送验证码
func (h *UserInfoHandler) SendSmsCode(ctx *gin.Context) {
	var req req.SmsCodeSendReq
	if err := ctx.ShouldBind(&req); err != nil {
		JsonBack(ctx, err.Error(), -2, nil)
		return
	}
	err := h.userInfoService.SendSmsCode(req)
	if err != nil {
		JsonBack(ctx, err.Error(), -1, nil)
		return
	}
	JsonBack(ctx, "发送成功", 0, nil)
}

// 注册
func (h *UserInfoHandler) Register(ctx *gin.Context) {
	var req req.RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		JsonBack(ctx, err.Error(), -2, nil)
		return
	}
	err := h.userInfoService.Register(req)
	if err != nil {
		JsonBack(ctx, err.Error(), -1, nil)
		return
	}
	JsonBack(ctx, "注册成功", 0, nil)
}
