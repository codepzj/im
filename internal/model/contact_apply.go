package model

import (
	"time"

	"gorm.io/gorm"
)

// 申请加好友,群聊
type ContactApply struct {
	ID          uint   `gorm:"primarykey"`
	UserID      string `gorm:"column:user_id"`       // 申请人uuid
	ContactID   string `gorm:"column:contact_id"`    // 被申请的uuid【用户/群聊】
	ContactType int8   `gorm:"column:contact_type"`  // 申请类型 0 用户 1群聊
	Status      int8   `gorm:"column:status"`        // 状态 0 申请中 1 申请通过 2 拒绝 3 拉黑
	Message     string `gorm:"column:message"`       // 申请信息
	LastApplyAt string `gorm:"column:last_apply_at"` // 最后申请时间
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
