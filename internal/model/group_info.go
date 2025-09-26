package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// 群成员信息
type GroupInfo struct {
	ID          uint            `gorm:"primarykey"`
	Uuid        string          `gorm:"column:uuid"`                   // 唯一标识
	Name        string          `gorm:"column:name"`                   // 群聊名称
	Notice      string          `gorm:"column:notice"`                 // 群聊公告
	Avatar      string          `gorm:"column:avatar"`                 // 群聊头像
	Members     json.RawMessage `gorm:"column:members"`                // 用户uuid
	MemberCount int             `gorm:"column:member_count;default:1"` // 群聊人数【默认群主1人】
	OwnerID     string          `gorm:"column:owner_id"`               // 群主uuid
	AddMode     int8            `gorm:"column:add_mode"`               // 0 直接 1 审核
	Status      int8            `gorm:"column:status"`                 // 0 正常 1 禁用 2 解散
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
