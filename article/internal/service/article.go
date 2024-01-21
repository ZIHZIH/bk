package service

import (
	"bk/article/api/pb"
	"bk/article/internal/dao"
	"bk/article/utils/kafka"
	"context"
)

type ArticleService struct {
	pb.UnimplementedArticleServiceServer
	ArticleMongodbDao *dao.ArticleMongodbDao
	Producer          *kafka.Producer
}

func (A *ArticleService) CreateArticle(ctx context.Context, request *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	res, err := A.ArticleMongodbDao.ArticleInsertOne(ctx, request.Article)
	if err != nil {
		return nil, err
	}

	return &pb.CreateArticleResponse{Article: res}, nil
}

func (A *ArticleService) DeleteArticle(ctx context.Context, request *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	if err := A.ArticleMongodbDao.ArticleDeleteOne(ctx, request.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteArticleResponse{Deleted: true}, nil
}

func (A *ArticleService) GetArticle(ctx context.Context, request *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	// 写日志
	A.Producer.Write(ctx, "GetArticle be called")
	res, err := A.ArticleMongodbDao.ArticleFindOne(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleResponse{Id: res}, nil
}

func (A *ArticleService) UpdateArticle(ctx context.Context, request *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	res, err := A.ArticleMongodbDao.ArticleUpdateOne(ctx, request.Article)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateArticleResponse{Count: int32(res)}, nil
}

func (A *ArticleService) ListArticle(ctx context.Context, request *pb.ListArticleRequest) (*pb.ListArticleResponse, error) {
	res, err := A.ArticleMongodbDao.ArticleListFind(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListArticleResponse{Articles: res}, nil
}

func NewArticleService(ArticleMongodbDao *dao.ArticleMongodbDao, Producer *kafka.Producer) *ArticleService {
	return &ArticleService{
		ArticleMongodbDao: ArticleMongodbDao,
		Producer:          Producer,
	}
}
