package internal

import (
	"fmt"
	"net/http"
	"wzh/controller"
)

// UserLogin 用户登陆
func UserLogin(res http.ResponseWriter, req *http.Request) {
	fmt.Println("用户登陆")
	defer func() { _ = req.Body.Close() }()
	// 从req中提取所需要的信息
	err := req.ParseForm()
	username := req.PostForm.Get("username")
	password := req.PostForm.Get("password")
	// 根据提供的用户名查询数据库
	pw, err := controller.GetUser(username)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}
	// 密码是否正确
	if pw != password {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("密码错误"))
		return
	}
	// 填写响应消息
	res.Write(httpResponseSuccessMessage)
}

// UserRegister 用户注册
func UserRegister(res http.ResponseWriter, req *http.Request) {
	fmt.Println("用户注册")
	defer func() { _ = req.Body.Close() }()
	req.ParseForm()
	username := req.PostForm.Get("username")
	password := req.PostForm.Get("password")
	// 根据用户查询数据库，判断用户是否存在
	_, err := controller.GetUser(username)
	if err == nil {
		res.Write([]byte("用户已经存在"))
		return
	}
	controller.CreateUser(username, password)
	// 填写响应消息
	res.Write(httpResponseSuccessMessage)
}
