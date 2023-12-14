package main

import (
	"fmt"
	"net/http"
	"wzh/internal"
)

func main() {
	http.HandleFunc("/index", func(res http.ResponseWriter, req *http.Request) { res.Write([]byte("HELLO WELCOME WZH WEB")) })
	http.HandleFunc("/user/login", internal.UserLogin)
	http.HandleFunc("/user/register", internal.UserRegister)
	http.HandleFunc("/article", internal.ArticleProcess)

	// 启动服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start failed,error is", err)
	}
}
