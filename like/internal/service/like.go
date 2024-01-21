package service

import (
	"bk/like/api/pb"
	"bk/like/internal/dao"
	model "bk/like/internal/dao/models"
	"bk/like/internal/dto"
	"context"
)

type LikeService struct {
	pb.UnimplementedLikeServiceServer
	LikeDao *dao.LikeDao
}

func (L *LikeService) CreateLike(ctx context.Context, req *pb.CreateLikeRequest) (*pb.CreateLikeResponse, error) {
	res, err := L.LikeDao.CreatLike(ctx, &model.Like{ArticleId: int(req.ArticleId), LikerId: int(req.LikerId)})
	if err != nil {
		return nil, err
	}

	likeDto, err := dto.Like(ctx, res)
	if err != nil {
		return nil, err
	}

	return &pb.CreateLikeResponse{Like: likeDto}, nil
}

func (L *LikeService) GetLikeByArticleId(ctx context.Context, req *pb.GetLikeByArticleIdRequest) (*pb.GetLikeByArticleIdResponse, error) {
	res, err := L.LikeDao.GetLikeByArticleId(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	likes := make([]*pb.Like, 0)
	for _, record := range res {
		like, err := dto.Like(ctx, record)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return &pb.GetLikeByArticleIdResponse{Likes: likes}, nil
}

func NewLikeService(LikeDao *dao.LikeDao) *LikeService {
	return &LikeService{
		LikeDao: LikeDao,
	}
}
