package service

import (
	"context"
	"errors"
	"fmt"
	"im/internal/config"
	"im/internal/dao"
	"im/internal/dto/req"
	"im/internal/global/constrants"
	"im/internal/global/enums"
	"im/internal/model"
	"im/pkg"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserInfoService struct {
	conf *config.Config
	rdb  *redis.Client
	dao  *dao.UserInfoDao
}

func NewUserInfoService(conf *config.Config, rdb *redis.Client, dao *dao.UserInfoDao) *UserInfoService {
	return &UserInfoService{
		conf: conf,
		rdb:  rdb,
		dao:  dao,
	}
}

// Login 登录
func (s *UserInfoService) Login(req req.LoginReq) error {
	userInfo, err := s.dao.FindOne(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("phone=?", req.Phone)
	})
	// 用户不存在
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return constrants.ErrUserNotFound
	}
	if userInfo == nil {
		return constrants.ErrUserLoadFailed
	}
	// 用户密码错误
	if userInfo.Password != req.Password {
		return constrants.ErrUserPasswordError
	}
	return nil
}

// LoginWithSms 验证码登录
func (s *UserInfoService) LoginWithSms(req req.LoginWithSmsReq) error {
	key := fmt.Sprintf("auth_code_%s", req.Phone)
	code, err := s.rdb.Get(context.Background(), key).Result()
	// 未命中缓存
	if err != nil || code == "" {
		return constrants.RandomCodeExpired
	}

	// 校验验证码是否匹配
	if code != req.PhoneCode {
		return constrants.ErrRandomCodeMatch
	}
	
	// 查询数据库中是否有这条记录
	_, err = s.dao.FindOne(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("phone=?", req.Phone)
	})
	if err != nil {
		return constrants.ErrUserNotFound
	}
	return nil
}

// 发送验证码
/*
查看redis中是否有对应的key
- 有，返回错误说明验证码未过期
- 无，则发送
*/
func (s *UserInfoService) SendSmsCode(req req.SmsCodeSendReq) error {
	key := fmt.Sprintf("auth_code_%s", req.Phone)

	// 未命中缓存
	if _, err := s.rdb.Get(context.Background(), key).Result(); err != nil {
		// 生成随机6位验证码, 有效时间为10分钟
		phoneCode := pkg.GenRandomCode(6)
		s.rdb.SetNX(context.Background(), key, phoneCode, time.Duration(s.conf.SmsExpireTime)*time.Minute)

		tecentSms := s.conf.TecentSms
		if err := pkg.SendTecentSmsCode(tecentSms.AccessKeyId, tecentSms.AccessKeySecret, tecentSms.SmsSdkAppId, tecentSms.SignName, tecentSms.TemplateId, phoneCode, req.Phone); err != nil {
			// 打日志
			go func() {
				slog.Error("发送验证码失败", "phone", req.Phone, "err", err.Error())
			}()
			return constrants.ErrSendRandomCode
		}
		return nil
	}
	// 命中缓存则不发送
	return constrants.RandomCodeNotExpired
}

// 注册
/*
1. 校验验证码
2. 校验手机号是否被注册
3. 创建用户
*/
func (s *UserInfoService) Register(registerReq req.RegisterReq) error {
	key := fmt.Sprintf("auth_code_%s", registerReq.Phone)
	val, err := s.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return constrants.RandomCodeExpired // 验证码已过期
	}
	if val != registerReq.PhoneCode {
		return constrants.ErrRandomCodeMatch // 验证码错误
	}
	userInfo, err := s.dao.FindOne(func(tx *gorm.DB) *gorm.DB {
		return tx.Unscoped().Where("phone=?", registerReq.Phone)
	})
	fmt.Println("userInfo", userInfo)
	// 校验手机号是否被注册
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err := s.dao.Create(&model.UserInfo{
			Uuid:      uuid.NewString(),
			NickName:  registerReq.NickName,
			Password:  registerReq.Password,
			Phone:     registerReq.Phone,
			Age:       18,
			Sex:       int8(enums.Male),
			Email:     "2363435714@qq.com",
			Avatar:    "https://element-plus.org/images/element-plus-logo.svg",
			Signature: "这个人很懒，什么都没留下",
			IsAdmin:   0,
			Status:    int8(enums.UserStatusNormal),
		})
		if err != nil {
			return err
		}
		return nil // 注册成功
	}
	return constrants.ErrPhoneHasRegister
}
