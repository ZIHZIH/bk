package controller

import (
	"context"
	"wzh/dal/model"
	"wzh/infra"
)

// GetArticle 根据id查询文章记录
func GetArticle(ctx context.Context, id int) (*model.Article, error) {
	article := new(model.Article)
	if ret := infra.DB.First(article, id); ret.Error != nil {
		return nil, ret.Error
	}
	return article, nil
}

// CreatArticle 创建新的文章记录
func CreatArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
	if ret := infra.DB.Create(article); ret.Error != nil {
		return nil, ret.Error
	}

	return article, nil
}

// DeleteArticle 根据文章id对文章记录进行删除
func DeleteArticle(ctx context.Context, id int) error {
	if ret := infra.DB.Delete(&model.Article{}, id); ret.Error != nil {
		return ret.Error
	}
	return nil
}

// UpdateArticle 对文章进行更新
func UpdateArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
	if err := infra.DB.Model(article).Updates(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

// ListArticle 获取所有文章列表
func ListArticle(ctx context.Context) ([]*model.Article, error) {
	result := make([]*model.Article, 0)
	if err := infra.DB.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
