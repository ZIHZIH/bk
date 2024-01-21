package dto

import (
	"bk/comment/api/pb"
	model "bk/comment/internal/dao/models"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Comment(ctx context.Context, comment *model.Comment) (*pb.Comment, error) {
	return &pb.Comment{
		Id:            int32(comment.Id),
		ArticleId:     int32(comment.ArticleId),
		CommentatorId: int32(comment.CommentatorId),
		Content:       comment.Content,
		CreateTime:    timestamppb.New(comment.CreatedAt),
		UpdateTime:    timestamppb.New(comment.UpdatedAt),
	}, nil
}
