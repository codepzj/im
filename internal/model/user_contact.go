package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户联系表【用户与用户或者群聊之间的联系】
type UserContact struct {
	ID          uint `gorm:"primarykey"`
	UserID      uint `gorm:"column:user_id"`      // 用户的uuid
	ContactID   uint `gorm:"column:contact_id"`   // 用户或群聊的uuid
	ContactType int8 `gorm:"column:contact_type"` // 0 用户 1 群聊
	Status      int8 `gorm:"column:status"`       // 0 正常 1 拉黑 2 被拉黑 3 删除好友 4 被删除好友 5 被禁言 6 退出群聊 7 被踢出群聊
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
