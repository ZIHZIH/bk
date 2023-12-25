package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		timeConsuming := time.Since(start)
		fmt.Println("这次处理耗时为", timeConsuming)
	}
}

func InitGinRouter(r *gin.Engine) {
	r.Use(MiddleWare())
	{
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "wzh's web")
		})
		r.POST("/user/login", userLogin)
		r.POST("/user/register", userRegister)
		r.GET("/getArticle", articleGet)
		r.GET("/listArticle", listArticle)
		r.POST("/createArticle", articleCreat)
		r.PUT("/updateArticle", articleUpdate)
		r.DELETE("/deleteArticle", articleDelete)
	}
}
