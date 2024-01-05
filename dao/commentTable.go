package dao

import (
	"context"
	"wzh/infra"
	"wzh/model"
)

// CreatComment 创建点赞记录
func CreatComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	if err := infra.DB.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentByArticleId 通过文章id获得该文章的评论信息
func GetCommentByArticleId(ctx context.Context, articleId int) ([]*model.Comment, error) {
	var res []*model.Comment
	if err := infra.DB.Where("article_id = ?", articleId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
