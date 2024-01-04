package controller

import (
	"context"
	"wzh/dal/model"
	"wzh/infra"
)

// CreatLike 创建点赞记录
func CreatLike(ctx context.Context, like *model.Like) (*model.Like, error) {
	if err := infra.DB.Create(like).Error; err != nil {
		return nil, err
	}

	return like, nil
}
