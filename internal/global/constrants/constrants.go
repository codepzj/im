package constrants

import "errors"

// 用户
var (
	// 用户不存在
	ErrUserNotFound = errors.New("用户不存在")
	// 用户信息异常
	ErrUserLoadFailed = errors.New("用户信息异常")
	// 手机号已经被注册
	ErrPhoneHasRegister = errors.New("手机号已经被注册")
	// 用户密码错误
	ErrUserPasswordError = errors.New("用户密码错误")
)

// 验证码
var (
	// 验证码未过期
	RandomCodeNotExpired = errors.New("验证码未过期, 请输入发送过的验证码")
	// 验证码已过期
	RandomCodeExpired = errors.New("验证码已过期")
	// 发送验证码失败
	ErrSendRandomCode = errors.New("发送验证码失败")
	// 验证码错误
	ErrRandomCodeMatch = errors.New("验证码错误")
)

// 数据库
var (
	// 创建记录失败
	ErrCreatedFailed = errors.New("创建记录失败")
)
