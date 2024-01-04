package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wzh/infra"
	"wzh/router"
)

func main() {
	r := gin.Default()
	router.Init(r)

	if err := infra.Init(); err != nil {
		fmt.Println(err)
		return
	}

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
