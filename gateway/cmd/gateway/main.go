package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wzh/gateway/config"
	"wzh/gateway/controller"
	"wzh/gateway/utils"
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
