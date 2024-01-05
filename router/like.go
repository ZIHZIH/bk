package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wzh/dao"
	"wzh/model"
)

// LikeArticle 评论文章
func LikeArticle(c *gin.Context) {
	temp := new(model.Like)
	err := c.ShouldBind(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	result, err := dao.CreatLike(c, temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
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

	result, err := dao.GetLikeByArticleId(c, queryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, len(result))
}
