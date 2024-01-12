package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wzh/gateway/utils/auth_jwt"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		timeConsuming := time.Since(start)
		fmt.Println("这次处理耗时为", timeConsuming)
	}
}

func Init(r *gin.Engine) {
	r.Use(MiddleWare())
	{
		r.GET("/", func(c *gin.Context) {
			//c.String(http.StatusOK, "wzh's web")
		})
		r.POST("/user/login", UserLogin)
		r.POST("/user/register", UserRegister)
		r.GET("/getArticle", ArticleGet)
		r.GET("/listArticle", ListArticle)
		r.POST("/createArticle", ArticleCreat)
		r.PUT("/updateArticle", ArticleUpdate)
		r.DELETE("/deleteArticle", ArticleDelete)
		r.POST("/commentArticle", CommentArticle)
		r.GET("/getArticleComment", CommentGetByArticleID)
		r.POST("/likeArticle", LikeArticle)
		r.GET("/getArticleLike", LikeGetByArticleID)
		r.GET("/import", auth_jwt.AuthMiddleware(), func(c *gin.Context) {
			value, exists := c.Get("UserId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "没找到用户id jwt有问题")
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "访问到了最最重要的数据",
				"You id":  value,
			})
		})
	}
}
