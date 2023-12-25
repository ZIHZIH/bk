package controller

import (
	"errors"
	"time"
)

var userTotal int

// 第一个key为id，模拟数据库保存相应的用户数据
var userMap map[int]*UserRecord

// k：手机号 v：userid
var PhoneNumberMap map[string]int

type UserRecord struct {
	Id          int       `json:"id"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	Password    string    `json:"password" form:"password"`
	Identity    string    `json:"identity" form:"identity"`
	IpPosition  string    `json:"id_position" form:"id_position"`
	CreateTime  time.Time `json:"create_time" form:"create_time"`
}

func init() {
	userMap = make(map[int]*UserRecord)
	PhoneNumberMap = make(map[string]int)
}

// GetUser 根据username查询密码
func GetUser(phoneNumber string) (*UserRecord, error) {
	if v, ok := PhoneNumberMap[phoneNumber]; ok {
		if record, ok := userMap[v]; ok {
			return record, nil
		}
	}

	return nil, errors.New("user is not exist")
}

// CreateUser 创建一个新的用户
func CreateUser(record *UserRecord) (*UserRecord, error) {
	// 判断用户是否已存在
	if _, ok := PhoneNumberMap[record.PhoneNumber]; ok {
		return nil, errors.New("用户已经存在")
	}

	userTotal++
	record.Id = userTotal
	record.CreateTime = time.Now()

	// 入库记录
	userMap[userTotal] = record
	PhoneNumberMap[record.PhoneNumber] = userTotal
	return record, nil
}
