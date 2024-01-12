package dao

//
//import (
//	"context"
//	"testing"
//	"wzh/ability/infra"
//	"wzh/article/internal/server"
//	"wzh/comment/internal/service"
//	dao2 "wzh/like/internal/dao"
//	"wzh/model"
//	"wzh/user/internal/dao"
//	"wzh/user/internal/dao/models"
//)
//
//func TestGetArticle(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := server.ArticleMongodbD.ArticleFindOne(context.Background(), "659d03cde9c04750ae6fac13")
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestCreatArticle(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := server.ArticleMongodbD.ArticleInsertOne(context.Background(), `{
//    "author_id":99999,
//    "title":"zzzzzzzzz article",
//    "content":"szszszszszsz",
//    "label":"night",
//    "status":3
//}`)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestUpdateArticle(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := server.ArticleMongodbD.ArticleUpdateOne(context.Background(), `{
//    "_id": "659cbf76dd2bb40b66dad5aa",
//    "author_id": 222222,
//    "title": "qmqmqmqmqmqmqmq article",
//    "content": "jjjjjjjjjjj",
//    "label": "one",
//    "Status": 3,
//    "ID": 4,
//    "CreatedAt": "2024-01-03T21:18:46.61+08:00",
//    "UpdatedAt": "2024-01-03T21:18:46.61+08:00",
//    "DeletedAt": null
//}`)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestDeleteArticlee(t *testing.T) {
//	infra.Init()
//	server.Init()
//	err := server.ArticleMongodbD.ArticleDeleteOne(context.Background(), "")
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestListArticle(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := server.ArticleMongodbD.ArticleListFind(context.Background())
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestGetUser(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := dao.GetUser(context.Background(), "19157692290")
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestCreateUser(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := dao.CreateUser(context.Background(), &models.User{
//		Id:          0,
//		Username:    "zxj",
//		PhoneNumber: "189189177",
//		Password:    "789780",
//		Avatar:      "w",
//		Identity:    "z",
//		IpPosition:  "sh",
//	})
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestGetLikeByArticleId(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := dao2.GetLikeByArticleId(context.Background(), 3)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestCreatLike(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := dao2.CreatLike(context.Background(), &model.Like{
//		ArticleId: 7,
//		LikerId:   3,
//	})
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestCreatComment(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := service.CreatComment(context.Background(), &model.Comment{
//		ArticleId:     7,
//		CommentatorId: 3,
//		Content:       "qweqwe",
//	})
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
//
//func TestGetCommentByArticleId(t *testing.T) {
//	infra.Init()
//	server.Init()
//	_, err := service.GetCommentByArticleId(context.Background(), 3)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//}
