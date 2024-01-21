package dto

import (
	"bk/user/api/pb"
	"bk/user/internal/dao/models"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func User(ctx context.Context, user *models.User) (*pb.User, error) {
	return &pb.User{
		Id:          int32(user.Id),
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		Avatar:      user.Avatar,
		Identity:    user.Identity,
		IpPosition:  user.IpPosition,
		CreateTime:  timestamppb.New(user.CreatedAt),
		UpdateTime:  timestamppb.New(user.UpdatedAt),
	}, nil
}
