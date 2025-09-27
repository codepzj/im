package dao

import (
	"im/internal/global/constrants"
	"im/internal/model"

	"gorm.io/gorm"
)

type UserInfoDao struct {
	db *gorm.DB
}

func NewUserInfoDao(db *gorm.DB) *UserInfoDao {
	return &UserInfoDao{
		db: db,
	}
}

func (d *UserInfoDao) FindOne(filter func(tx *gorm.DB) *gorm.DB) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	err := d.db.Model(&model.UserInfo{}).Scopes(filter).First(&userInfo).Error
	return &userInfo, err
}

func (d *UserInfoDao) Create(user *model.UserInfo) error {
	res := d.db.Model(&model.UserInfo{}).Create(user)
	if res.RowsAffected == 0 {
		return constrants.ErrCreatedFailed
	}
	return res.Error
}
