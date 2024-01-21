package service

import (
	"bk/user/api/pb"
	"bk/user/internal/dao"
	"bk/user/internal/dao/models"
	"bk/user/internal/dto"
	"context"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	UserDao *dao.UserDao
}

func (u *UserService) UserRegister(ctx context.Context, request *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	user, err := u.UserDao.CreateUser(ctx, &models.User{
		Username:    request.Username,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
		Avatar:      request.Avatar,
		Identity:    request.Identity,
		IpPosition:  request.IpPosition,
	})
	if err != nil {
		return nil, err
	}

	resp, err := dto.User(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserRegisterResponse{User: resp}, nil
}

func (u *UserService) GetUserByPhoneNumber(ctx context.Context, request *pb.GetUserByPhoneNumberRequest) (*pb.GetUserByPhoneNumberResponse, error) {
	res, err := u.UserDao.GetUserByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		return nil, err
	}

	resp, err := dto.User(ctx, res)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserByPhoneNumberResponse{User: resp}, nil
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{UserDao: userDao}
}
