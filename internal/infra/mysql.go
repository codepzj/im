package infra

import (
	"im/internal/config"
	"im/internal/model"
	"log"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewMySQL 初始化MySQL数据库连接
func NewMySQL(cfg *config.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(getLogOutput(cfg), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Duration(cfg.Mysql.SlowSqlThreshold) * time.Second, // 慢查询阈值，单位秒
			LogLevel:      parseLogLevel(cfg),                                      // 日志级别
		},
	)
	db, err := gorm.Open(mysql.Open(cfg.Mysql.Dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		slog.Error("连接Mysql失败", "error", err)
		panic(err.Error())
	}

	// 自动迁移数据库结构
	if err := db.AutoMigrate(&model.UserInfo{}, &model.GroupInfo{}, &model.UserContact{},
		&model.ContactApply{}, &model.Session{}, &model.Message{},
	); err != nil {
		slog.Error("自动迁移数据库结构失败", "error", err)
	}

	return db
}

func getLogOutput(cfg *config.Config) *os.File {
	switch cfg.Mysql.LogType {
	case 0:
		return os.Stdout
	case 1:
		f, err := os.OpenFile(cfg.Mysql.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			slog.Error("打开日志文件失败", "error", err)
		}
		return f
	}
	return os.Stdout
}

func parseLogLevel(cfg *config.Config) logger.LogLevel {
	switch cfg.Mysql.LogLevel {
	case "silent":
		return logger.Silent
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	}
	return logger.Info
}
