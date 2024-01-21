package main

import (
	"bk/gateway/config"
	"bk/gateway/controller"
	"bk/gateway/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	config.Init()
	r := gin.Default()
	controller.Init(r)
	utils.NewUserServiceClient()
	utils.NewCommentServiceClient()
	utils.NewArticleServiceClient()
	utils.NewLikeServiceClient()
	err := r.Run(config.Config.AppConfig.Host + ":" + strconv.Itoa(config.Config.AppConfig.Port))
	if err != nil {
		panic(err)
	}
}
