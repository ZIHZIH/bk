package service

import (
	"bk/comment/api/pb"
	"bk/comment/internal/dao"
	model "bk/comment/internal/dao/models"
	"bk/comment/internal/dto"
	"context"
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

	resp, err := dto.Comment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentResponse{Comment: resp}, nil
}

func (c *CommentService) GetCommentByArticleId(ctx context.Context, request *pb.GetCommentByArticleIdRequest) (*pb.GetCommentByArticleIdResponse, error) {
	records, err := c.CommentDao.GetCommentByArticleId(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	comments := make([]*pb.Comment, 0)
	for _, comment := range records {
		res, err := dto.Comment(ctx, comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, res)
	}

	return &pb.GetCommentByArticleIdResponse{Comments: comments}, nil
}

func NewCommentService(CommentDao *dao.CommentDao) *CommentService {
	return &CommentService{CommentDao: CommentDao}
}
