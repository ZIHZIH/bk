package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"wzh/pkg/pb"
	"wzh/user/internal/dao"
	"wzh/user/internal/dao/models"
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

	return &pb.UserRegisterResponse{User: &pb.User{
		Id:          int32(user.Id),
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		Avatar:      user.Avatar,
		Identity:    user.Identity,
		IpPosition:  user.IpPosition,
		CreateTime:  timestamppb.New(user.CreatedAt),
		UpdateTime:  timestamppb.New(user.UpdatedAt),
	}}, nil
}

func (u *UserService) GetUserByPhoneNumber(ctx context.Context, request *pb.GetUserByPhoneNumberRequest) (*pb.GetUserByPhoneNumberResponse, error) {
	res, err := u.UserDao.GetUserByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByPhoneNumberResponse{User: &pb.User{
		Id:          int32(res.Id),
		Username:    res.Username,
		PhoneNumber: res.PhoneNumber,
		Password:    res.Password,
		Avatar:      res.Avatar,
		Identity:    res.Identity,
		IpPosition:  res.IpPosition,
		CreateTime:  timestamppb.New(res.CreatedAt),
		UpdateTime:  timestamppb.New(res.UpdatedAt),
	}}, nil
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{UserDao: userDao}
}
