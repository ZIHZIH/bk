package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wzh/controller"
)

// articleGet 文章的获取
func articleGet(c *gin.Context) {
	articleId := c.Query("id")
	if articleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dont have parameter：id"})
		return
	}

	recordId, err := strconv.Atoi(articleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.GetArticle(recordId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// articleUpdate 文章的更新
func articleUpdate(c *gin.Context) {
	temp := &controller.ArticleRecord{}
	err := c.ShouldBindJSON(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.UpdateArticle(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// articleDelete 文章的删除
func articleDelete(c *gin.Context) {
	articleId := c.Query("id")
	if articleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dont have parameter：id"})
		return
	}

	recordId, err := strconv.Atoi(articleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = controller.DeleteArticle(recordId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// articleCreat 文章的创建
func articleCreat(c *gin.Context) {
	temp := &controller.ArticleRecord{}
	err := c.ShouldBindJSON(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.CreatArticle(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// listArticle 列出所有文章
func listArticle(c *gin.Context) {
	resp, err := controller.ListArticle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
