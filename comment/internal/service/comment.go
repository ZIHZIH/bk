package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"wzh/comment/internal/dao"
	model "wzh/comment/internal/dao/models"
	"wzh/pkg/pb"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
	CommentDao *dao.CommentDao
}

func (c *CommentService) CreateComment(ctx context.Context, request *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	comment, err := c.CommentDao.CreatComment(ctx, &model.Comment{ArticleId: int(request.ArticleId), CommentatorId: int(request.CommentatorId), Content: request.Content})
	if err != nil {
		return nil, err
	}
	return &pb.CreateCommentResponse{Comment: &pb.Comment{
		Id:            int32(comment.Id),
		ArticleId:     int32(comment.ArticleId),
		CommentatorId: int32(comment.CommentatorId),
		Content:       comment.Content,
		CreateTime:    timestamppb.New(comment.CreatedAt),
		UpdateTime:    timestamppb.New(comment.UpdatedAt),
	}}, nil
}

func (c *CommentService) GetCommentByArticleId(ctx context.Context, request *pb.GetCommentByArticleIdRequest) (*pb.GetCommentByArticleIdResponse, error) {
	records, err := c.CommentDao.GetCommentByArticleId(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	comments := make([]*pb.Comment, 0)
	for _, comment := range records {
		temp := &pb.Comment{Id: int32(comment.Id), ArticleId: int32(comment.ArticleId), CommentatorId: int32(comment.Id), Content: comment.Content, CreateTime: timestamppb.New(comment.CreatedAt), UpdateTime: timestamppb.New(comment.UpdatedAt)}
		comments = append(comments, temp)
	}

	return &pb.GetCommentByArticleIdResponse{Comments: comments}, nil
}

func NewCommentService(CommentDao *dao.CommentDao) *CommentService {
	return &CommentService{CommentDao: CommentDao}
}
