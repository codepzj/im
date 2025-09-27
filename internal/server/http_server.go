package server

import (
	v1 "im/api/v1"
	"im/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var GE *gin.Engine

func InitRouter(conf *config.Config) {
	GE = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	GE.Use(cors.New(corsConfig))

	userInfoHandler := v1.NewUserInfoHandler(conf)

	{
		GE.GET("send/smsCode", userInfoHandler.GetSmsCode)
	}

	GE.Run(":9999")
}
