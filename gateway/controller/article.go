package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"wzh/gateway/utils"
	"wzh/pkg/pb"
)

// ArticleGet 文章的获取
func ArticleGet(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dont have parameter：id"})
		return
	}

	resp, err := utils.ArticleServiceClient.GetArticle(context.Background(), &pb.GetArticleRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Id)
}

// ArticleUpdate 文章的更新
func ArticleUpdate(c *gin.Context) {
	b := make([]byte, c.Request.ContentLength)
	_, err := c.Request.Body.Read(b)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := utils.ArticleServiceClient.UpdateArticle(context.Background(), &pb.UpdateArticleRequest{Article: string(b)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Count)
}

// ArticleDelete 文章的删除
func ArticleDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dont have parameter：id"})
		return
	}

	resp, err := utils.ArticleServiceClient.DeleteArticle(context.Background(), &pb.DeleteArticleRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Deleted)
}

// ArticleCreat 文章的创建
func ArticleCreat(c *gin.Context) {
	b := make([]byte, c.Request.ContentLength)
	_, err := c.Request.Body.Read(b)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := utils.ArticleServiceClient.CreateArticle(context.Background(), &pb.CreateArticleRequest{Article: string(b)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Article)
}

// ListArticle 列出所有文章
func ListArticle(c *gin.Context) {
	resp, err := utils.ArticleServiceClient.ListArticle(context.Background(), &pb.ListArticleRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Articles)
}
