package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wzh/controller"
	"wzh/dal/model"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	temp := new(model.User)
	err := c.ShouldBind(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	if temp.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "电话号码为空"})
		return
	}
	if temp.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户昵称为空"})
		return
	}

	resp, err := controller.CreateUser(c, temp)
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

	record, err := controller.GetUser(c, phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	password := c.PostForm("password")
	if record.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确"})
		return
	}

	c.JSON(http.StatusOK, "密码正确，登陆成功")
}
