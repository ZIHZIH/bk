package dao

import (
	"wzh/infra"
	"wzh/pkg/cache"
)

var ArticleD *ArticleDao

func Init() {
	ArticleD = NewArticleDao(infra.DB, cache.NewCache("wzh:bk:article:", infra.RedisDb))
}
