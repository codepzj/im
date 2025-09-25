package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户信息
type UserInfo struct {
	ID        uint   `gorm:"primarykey"`
	Uuid      string `gorm:"column:uuid"`        // 唯一标识
	NickName  string `gorm:"column:nickname"`    // 昵称
	Phone     string `gorm:"column:phone;index"` // 手机
	Age       int    `gorm:"column:age"`         // 年龄
	Sex       int    `gorm:"column:sex"`         // 性别 0 男 1 女 2 未知
	Email     string `gorm:"column:email"`       // 邮箱
	Avatar    string `gorm:"column:avatar"`      // 头像
	Signature string `gorm:"column:signature"`   // 个性签名
	Birthday  string `gorm:"column:birthday"`    // 生日
	IsAdmin   int8   `gorm:"column:is_admin"`    // 是否为管理员
	Status    int8   `gorm:"column:status"`      // 0 正常 1 禁用
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
