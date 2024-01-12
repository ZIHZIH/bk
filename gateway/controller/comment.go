package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wzh/gateway/utils"
	"wzh/pkg/pb"
)

// CommentArticle 评论文章
func CommentArticle(c *gin.Context) {
	req := new(pb.CreateCommentRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	result, err := utils.CommentServiceClient.CreateComment(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// CommentGetByArticleID 文章评论
func CommentGetByArticleID(c *gin.Context) {
	articleId := c.Query("article_id")
	if articleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dont have parameter：article_id"})
		return
	}

	queryId, err := strconv.Atoi(articleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := utils.CommentServiceClient.GetCommentByArticleId(context.Background(), &pb.GetCommentByArticleIdRequest{Id: int32(queryId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
