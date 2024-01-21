package dto

import (
	"bk/like/api/pb"
	model "bk/like/internal/dao/models"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Like(ctx context.Context, like *model.Like) (*pb.Like, error) {
	return &pb.Like{
		Id:         int32(like.Id),
		ArticleId:  int32(like.ArticleId),
		LikerId:    int32(like.LikerId),
		CreateTime: timestamppb.New(like.CreatedAt),
		UpdateTime: timestamppb.New(like.UpdatedAt),
	}, nil
}
