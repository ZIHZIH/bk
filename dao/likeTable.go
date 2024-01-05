package dao

import (
	"context"
	"wzh/infra"
	"wzh/model"
)

// CreatLike 创建点赞记录
func CreatLike(ctx context.Context, like *model.Like) (*model.Like, error) {
	if err := infra.DB.Create(like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

// GetLikeByArticleId 通过文章id获得该文章的点赞信息
func GetLikeByArticleId(ctx context.Context, articleId int) ([]*model.Like, error) {
	var res []*model.Like
	if err := infra.DB.Where("article_id = ?", articleId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
