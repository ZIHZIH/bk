package dao

import (
	"wzh/infra"
	"wzh/pkg/cache"
)

// var ArticleD *ArticleDao
var ArticleMongodbD *ArticleMongodbDao

func Init() {
	//ArticleD = NewArticleDao(infra.MysqlDB, cache.NewCache("wzh:bk:article:", infra.RedisDb))
	ArticleMongodbD = NewArticleMongodbDao(infra.Mongodb.Database("bk").Collection("article"), cache.NewCache("wzh:bk:article:", infra.RedisDb))
}
