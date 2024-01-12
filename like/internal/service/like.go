package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"wzh/like/internal/dao"
	model "wzh/like/internal/dao/models"
	"wzh/pkg/pb"
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

	return &pb.CreateLikeResponse{Like: &pb.Like{
		Id:         int32(res.Id),
		ArticleId:  int32(res.ArticleId),
		LikerId:    int32(res.LikerId),
		CreateTime: timestamppb.New(res.CreatedAt),
		UpdateTime: timestamppb.New(res.UpdatedAt),
	}}, nil
}

func (L *LikeService) GetLikeByArticleId(ctx context.Context, req *pb.GetLikeByArticleIdRequest) (*pb.GetLikeByArticleIdResponse, error) {
	res, err := L.LikeDao.GetLikeByArticleId(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	likes := make([]*pb.Like, 0)
	for _, record := range res {
		like := &pb.Like{
			Id:         int32(record.Id),
			ArticleId:  int32(record.ArticleId),
			LikerId:    int32(record.LikerId),
			CreateTime: timestamppb.New(record.CreatedAt),
			UpdateTime: timestamppb.New(record.UpdatedAt),
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
