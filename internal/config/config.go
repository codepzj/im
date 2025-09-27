package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	AppName string `mapstructure:"appName"` // 应用名称
	Host    string `mapstructure:"host"`    // 监听地址
	Port    int    `mapstructure:"port"`    // 监听端口
}

type MysqlConfig struct {
	Dsn              string `mapstructure:"dsn"`              // 数据库连接字符串
	LogType          int    `mapstructure:"logType"`          // 0 控制台 1 文件
	LogFile          string `mapstructure:"logFile"`          // 日志文件路径
	SlowSqlThreshold int    `mapstructure:"slowSqlThreshold"` // 慢查询阈值，单位秒
	LogLevel         string `mapstructure:"logLevel"`         // 日志级别 silent info warn error
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`     // Redis地址
	Port     string `mapstructure:"port"`     // Redis端口
	Password string `mapstructure:"password"` // Redis密码
	DB       int    `mapstructure:"db"`       // Redis数据库
}

type LogConfig struct {
	Level      string `mapstructure:"level"`      // 日志级别
	Filename   string `mapstructure:"filename"`   // 日志文件名
	MaxSize    int    `mapstructure:"maxSize"`    // 单个日志文件最大尺寸，单位MB
	MaxBackups int    `mapstructure:"maxBackups"` // 最大保留日志文件数量
	MaxAge     int    `mapstructure:"maxAge"`     // 最大保留天数
	Compress   bool   `mapstructure:"compress"`   // 是否压缩日志文件
}

type AliSmsConfig struct {
	AccessKeyId     string `mapstructure:"accessKeyId"`     // 阿里云accessKeyId
	AccessKeySecret string `mapstructure:"accessKeySecret"` // 阿里云accessKeySecret
}

type TecentSmsConfig struct {
	AccessKeyId     string `mapstructure:"accessKeyId"`     // 腾讯云accessKeyId
	AccessKeySecret string `mapstructure:"accessKeySecret"` // 腾讯云accessKeySecret
	SmsSdkAppId     string `mapstructure:"smsSdkAppId"`     // 腾讯云短信应用ID
	SignName        string `mapstructure:"signName"`        // 短信签名
	TemplateId      string `mapstructure:"templateId"`      // 短信模板ID
}

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Mysql     MysqlConfig     `mapstructure:"mysql"`
	Redis     RedisConfig     `mapstructure:"redis"`
	Log       LogConfig       `mapstructure:"log"`
	AliSms    AliSmsConfig    `mapstructure:"alisms"`
	TecentSms TecentSmsConfig `mapstructure:"tecentsms"`
}

func GetConfig() *Config {
	viper.SetConfigName("global")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败, %s", err.Error())
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("解析配置文件失败, %s", err.Error())
	}

	return &conf
}
