package model

import (
	"time"

	"gorm.io/gorm"
)

// 会话
/* 冗余头像和昵称可以保留历史状态, 与用户表解耦 */
type Session struct {
	ID          uint   `gorm:"primarykey"`
	Uuid        string `gorm:"column:uuid;type:varchar(36);uniqueIndex"` // 会话唯一标识
	SendID      string `gorm:"column:send_id"`          // 发送者的uuid
	ReceivedID  string `gorm:"column:received_id"`      // 接收者的uuid
	ReceiveName string `gorm:"column:receive_name"`     // 接收者的昵称
	Avatar      string `gorm:"column:avatar"`           // 会话头像
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
