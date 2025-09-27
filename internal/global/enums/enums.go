package enums

type (
	Sex        int8
	UserStatus int8
)

// 性别
const (
	Male   Sex = 0 // 男
	Female Sex = 1 // 女
)

// 用户状态
const (
	UserStatusNormal UserStatus = 0 // 正常
	UserStatusBan    UserStatus = 1 // 禁用
)
