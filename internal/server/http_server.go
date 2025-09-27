package server

import (
	v1 "im/api/v1"
	"im/internal/config"
	"im/internal/dao"
	"im/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var GE *gin.Engine

func InitRouter(conf *config.Config, db *gorm.DB, rdb *redis.Client) {
	GE = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	GE.Use(cors.New(corsConfig))

	/* 用户信息 */
	userInfoDao := dao.NewUserInfoDao(db)
	userInfoService := service.NewUserInfoService(conf, rdb, userInfoDao)
	userInfoHandler := v1.NewUserInfoHandler(userInfoService)

	{
		GE.GET("send/smsCode", userInfoHandler.SendSmsCode)       // 发送验证码
		GE.POST("login", userInfoHandler.Login)                   // 登录
		GE.POST("loginWithSms", userInfoHandler.LoginWithSmsCode) // 验证码登录
		GE.POST("register", userInfoHandler.Register)             // 注册
	}

	GE.Run(":9999")
}
