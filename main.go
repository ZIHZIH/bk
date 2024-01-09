package main

import (
	"github.com/gin-gonic/gin"
	"wzh/dao"
	"wzh/infra"
	"wzh/router"
)

func main() {
	infra.Init()
	dao.Init()

	r := gin.Default()
	router.Init(r)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
