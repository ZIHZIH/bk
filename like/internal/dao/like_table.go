package dao

import (
	"context"
	"log"
	"wzh/like/infra"
	model "wzh/like/internal/dao/models"
)

type LikeDao struct {
	logger *log.Logger
}

// CreatLike 创建点赞记录
func (dao *LikeDao) CreatLike(ctx context.Context, like *model.Like) (*model.Like, error) {
	if err := infra.MysqlDB.Create(like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

// GetLikeByArticleId 通过文章id获得该文章的点赞信息
func (dao *LikeDao) GetLikeByArticleId(ctx context.Context, articleId int) ([]*model.Like, error) {
	var res []*model.Like
	if err := infra.MysqlDB.Where("article_id = ?", articleId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func NewLikeDao(logger *log.Logger) *LikeDao {
	return &LikeDao{
		logger: logger,
	}
}
