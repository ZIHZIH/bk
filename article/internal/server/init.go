package server

//import (
//	"wzh/ability/cache"
//	"wzh/ability/infra"
//	"wzh/article/internal/dao"
//)
//
//// var ArticleD *ArticleDao
//var ArticleMongodbD *dao.ArticleMongodbDao
//
//func Init() {
//	//ArticleD = NewArticleDao(infra.MysqlDB, cache.NewCache("wzh:bk:article:", infra.RedisDb))
//	ArticleMongodbD = dao.NewArticleMongodbDao(infra.Mongodb.Database("bk").Collection("article"), cache.NewCache("wzh:bk:article:", infra.RedisDb))
//}
