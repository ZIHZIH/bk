package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wzh/controller"
)

// userRegister 用户注册
func userRegister(c *gin.Context) {
	temp := &controller.UserRecord{}
	err := c.ShouldBind(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	resp, err := controller.CreateUser(temp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// userLogin 用户登陆
func userLogin(c *gin.Context) {
	phoneNumber := c.PostForm("phone_number")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "电话号码为空"})
		return
	}

	record, err := controller.GetUser(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	password := c.PostForm("password")
	if record.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确"})
		return
	}

	c.JSON(http.StatusOK, record)
}
