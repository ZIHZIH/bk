package dao

import (
	"context"
	"wzh/infra"
	"wzh/model"
)

// GetUser 根据phoneNumber查询密码
func GetUser(ctx context.Context, phoneNumber string) (*model.User, error) {
	res := new(model.User)
	if err := infra.DB.Where("phone_number = ?", phoneNumber).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser 创建一个新的用户
func CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := infra.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
