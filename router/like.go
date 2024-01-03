package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wzh/controller"
	"wzh/dal/model"
)

// LikeArticle 评论文章
func LikeArticle(c *gin.Context) {
	temp := new(model.Like)
	err := c.ShouldBind(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	result, err := controller.CreatLike(c, temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
