package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wzh/controller"
	"wzh/dal/model"
)

// CommentArticle 评论文章
func CommentArticle(c *gin.Context) {
	temp := new(model.Comment)
	err := c.ShouldBind(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	result, err := controller.CreatComment(c, temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
