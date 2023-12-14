package controller

import (
	"errors"
)

var userMap map[string]string

func init() {
	userMap = make(map[string]string)
}

// GetUser 根据username查询密码
func GetUser(username string) (string, error) {
	if v, ok := userMap[username]; ok {
		return v, nil
	}

	return "", errors.New("user is not exist")
}

// CreateUser 创建一个新的用户
func CreateUser(username string, password string) bool {
	userMap[username] = password
	return true
}
