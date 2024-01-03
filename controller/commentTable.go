package controller

import (
	"context"
	"wzh/dal/model"
	"wzh/infra"
)

// CreatComment 创建点赞记录
func CreatComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	if err := infra.DB.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}
