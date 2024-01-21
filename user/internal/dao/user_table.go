package dao

import (
	"bk/user/infra"
	"bk/user/internal/dao/models"
	"context"
	"log"
)

type UserDao struct {
	Logger *log.Logger
}

// GetUser 根据phoneNumber查询密码
func (dao *UserDao) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	res := new(models.User)
	if err := infra.MysqlDB.Where("phone_number = ?", phoneNumber).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser 创建一个新的用户
func (dao *UserDao) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := infra.MysqlDB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserDao(logger *log.Logger) *UserDao {
	return &UserDao{Logger: logger}
}
