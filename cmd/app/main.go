package main

import (
	"im/internal/config"
	"im/internal/infra"
	"im/internal/server"
	"log"
)

func main() {
	conf := config.GetConfig()
	infra.InitLogger(conf) // 初始化日志

	_ = infra.NewMySQL(conf) // 数据库对象
	_ = infra.NewRedis(conf) // redis对象

	log.Println(conf)

	server.InitRouter(conf)
}
