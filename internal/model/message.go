package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID         uint   `gorm:"primarykey"`
	Uuid       string `gorm:"column:uuid;type:varchar(36);uniqueIndex"` // 唯一标识
	SessionID  string `gorm:"column:session_id"`                        // 会话ID
	SendType   int8   `gorm:"column:send_type"`                         // 发送类型 0 文本 1 文件 2 视频
	Content    string `gorm:"column:content"`                           // 内容
	Url        string `gorm:"column:url"`                               // 链接
	SendID     string `gorm:"column:send_id"`                           // 发送者的uuid
	SendName   string `gorm:"column:send_name"`                         // 发送者的昵称
	SendAvatar string `gorm:"column:send_avatar"`                       // 发送者的头像
	ReceivedID string `gorm:"column:received_id"`                       // 接收者的uuid
	FileName   string `gorm:"column:file_name"`                         // 文件名
	FileType   string `gorm:"column:file_type"`                         // 文件类型
	FileSize   string `gorm:"column:file_size"`                         // 文件大小
	Status     int8   `gorm:"column:status"`                            // 发送状态 0 为发送 1 已发送
	AVData     string `gorm:"column:av_data"`                           // 通话传递数据
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
