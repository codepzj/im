package main

import (
	"fmt"
	"im/internal/config"
	"im/internal/infra"
)

func main() {
	conf := config.GetConfig()
	infra.InitLogger(conf) // 初始化日志

	db := infra.NewMySQL(conf) // 初始化数据库
	fmt.Println(db)
	fmt.Println("Logger initialized")
}
