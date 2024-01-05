package dao

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"strconv"
	"time"
	"wzh/infra"
	"wzh/logger"
	"wzh/model"
	"wzh/pkg/cache"
)

type ArticleDao struct {
	DB    *gorm.DB
	Cache *cache.Cache
}

// GetArticle 根据id查询文章记录
func (articleDao *ArticleDao) GetArticle(ctx context.Context, id int) (*model.Article, error) {
	// 访问redis查询缓存，只要返回err就去访问数据库
	value, err := articleDao.Cache.GetCache(ctx, strconv.Itoa(id))
	if err == nil {
		res := new(model.Article)
		err = json.Unmarshal([]byte(value), res)
		if err == nil {
			return res, nil
		}
		logger.Logger.Println(err)
	}

	article := new(model.Article)
	if ret := infra.DB.First(article, id); ret.Error != nil {
		return nil, ret.Error
	}
	// 设置redis缓存
	b, err := json.Marshal(article)
	if err == nil {
		articleDao.Cache.SetCache(ctx, strconv.Itoa(id), b, 30*time.Minute)
	}

	return article, nil
}

// CreatArticle 创建新的文章记录
func (articleDao *ArticleDao) CreatArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
	if ret := infra.DB.Create(article); ret.Error != nil {
		return nil, ret.Error
	}

	return article, nil
}

// DeleteArticle 根据文章id对文章记录进行删除
func (articleDao *ArticleDao) DeleteArticle(ctx context.Context, id int) error {
	if ret := infra.DB.Delete(&model.Article{}, id); ret.Error != nil {
		return ret.Error
	}
	// 删除缓存
	articleDao.Cache.DelCache(ctx, strconv.Itoa(id))
	return nil
}

// UpdateArticle 对文章进行更新
func (articleDao *ArticleDao) UpdateArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
	if err := infra.DB.Model(article).Updates(article).Error; err != nil {
		return nil, err
	}
	// 删除缓存
	articleDao.Cache.DelCache(ctx, strconv.Itoa(article.Id))
	return article, nil
}

// ListArticle 获取所有文章列表
func (articleDao *ArticleDao) ListArticle(ctx context.Context) ([]*model.Article, error) {
	result := make([]*model.Article, 0)
	if err := infra.DB.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func NewArticleDao(db *gorm.DB, cache *cache.Cache) *ArticleDao {
	return &ArticleDao{
		DB:    db,
		Cache: cache,
	}
}
