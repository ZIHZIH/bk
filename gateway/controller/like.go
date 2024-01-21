package controller

import (
	"bk/gateway/api/pb"
	"bk/gateway/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// LikeArticle 评论文章
func LikeArticle(c *gin.Context) {
	req := new(pb.CreateLikeRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	resp, err := utils.LikeServiceClient.CreateLike(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// LikeGetByArticleID 文章点赞数的获取
func LikeGetByArticleID(c *gin.Context) {
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

	result, err := utils.LikeServiceClient.GetLikeByArticleId(context.Background(), &pb.GetLikeByArticleIdRequest{Id: int32(queryId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, len(result.Likes))
}
