package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wzh/controller"
	"wzh/dal/model"
)

// ArticleGet 文章的获取
func ArticleGet(c *gin.Context) {
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

	result, err := controller.GetArticle(c, recordId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ArticleUpdate 文章的更新
func ArticleUpdate(c *gin.Context) {
	temp := new(model.Article)
	err := c.ShouldBindJSON(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.UpdateArticle(c, temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ArticleDelete 文章的删除
func ArticleDelete(c *gin.Context) {
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

	err = controller.DeleteArticle(c, recordId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// ArticleCreat 文章的创建
func ArticleCreat(c *gin.Context) {
	temp := new(model.Article)
	err := c.ShouldBindJSON(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := controller.CreatArticle(c, temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ListArticle 列出所有文章
func ListArticle(c *gin.Context) {
	resp, err := controller.ListArticle(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
