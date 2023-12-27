package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wzh/internal"
)

func main() {
	r := gin.Default()
	internal.InitGinRouter(r)
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("c1test")
}
