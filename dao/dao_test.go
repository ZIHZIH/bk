package dao

import (
	"context"
	"testing"
	"wzh/infra"
	"wzh/logger"
	"wzh/model"
)

func TestGetArticle(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := ArticleD.GetArticle(context.Background(), 4)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCreatArticle(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := ArticleD.CreatArticle(context.Background(), &model.Article{AuthorId: 6, Title: "zzzz", Content: "szzzz", Status: 4})
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestUpdateArticle(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := ArticleD.UpdateArticle(context.Background(), &model.Article{Id: 100, AuthorId: 6})
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDeleteArticlee(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	err := ArticleD.DeleteArticle(context.Background(), 3)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestListArticle(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := ArticleD.ListArticle(context.Background())
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetUser(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := GetUser(context.Background(), "19157692290")
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCreateUser(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := CreateUser(context.Background(), &model.User{
		Id:          0,
		Username:    "zxj",
		PhoneNumber: "189189177",
		Password:    "789780",
		Avatar:      "w",
		Identity:    "z",
		IpPosition:  "sh",
	})
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetLikeByArticleId(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := GetLikeByArticleId(context.Background(), 3)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCreatLike(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := CreatLike(context.Background(), &model.Like{
		ArticleId: 7,
		LikerId:   3,
	})
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCreatComment(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := CreatComment(context.Background(), &model.Comment{
		ArticleId:     7,
		CommentatorId: 3,
		Content:       "qweqwe",
	})
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetCommentByArticleId(t *testing.T) {
	infra.Init()
	logger.Init()
	Init()
	_, err := GetCommentByArticleId(context.Background(), 3)
	if err != nil {
		t.Errorf("%s", err)
	}
}
