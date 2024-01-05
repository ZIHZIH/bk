package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wzh/dao"
	"wzh/model"
)

// CommentArticle 评论文章
func CommentArticle(c *gin.Context) {
	temp := new(model.Comment)
	err := c.ShouldBind(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	result, err := dao.CreatComment(c, temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// CommentGetByArticleID 文章评论数的获取
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

	result, err := dao.GetCommentByArticleId(c, queryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, len(result))
}
