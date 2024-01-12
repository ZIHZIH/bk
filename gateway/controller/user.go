package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"wzh/gateway/utils"
	"wzh/gateway/utils/auth_jwt"
	"wzh/pkg/pb"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	req := new(pb.UserRegisterRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	if req.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "电话号码为空"})
		return
	}
	if req.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户昵称为空"})
		return
	}

	resp, err := utils.UserServiceClient.UserRegister(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UserLogin 用户登陆
func UserLogin(c *gin.Context) {
	phoneNumber := c.PostForm("phone_number")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "电话号码为空"})
		return
	}

	resp, err := utils.UserServiceClient.GetUserByPhoneNumber(context.Background(), &pb.GetUserByPhoneNumberRequest{PhoneNumber: phoneNumber})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	password := c.PostForm("password")
	if resp.User.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确"})
		return
	}

	token, err := auth_jwt.GenToken(int(resp.User.Id), resp.User.Username, resp.User.PhoneNumber, resp.User.Avatar, resp.User.Identity, resp.User.IpPosition)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
