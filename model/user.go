package model

import "gorm.io/gorm"

// 用户表
type User struct {
	Id          int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:主键ID" json:"id" form:"id"`
	Username    string `gorm:"column:username;type:varchar(50);unique;comment:用户昵称;NOT NULL" json:"username" form:"username"`
	PhoneNumber string `gorm:"column:phone_number;type:char(11);unique;comment:手机号;NOT NULL" json:"phone_number" form:"phone_number"`
	Password    string `gorm:"column:password;type:varchar(16);comment:密码;NOT NULL" json:"password" form:"password"`
	Avatar      string `gorm:"column:avatar;type:blob;comment:头像" json:"avatar" form:"avatar"`
	Identity    string `gorm:"column:identity;type:varchar(50);comment:身份" json:"identity" form:"identity"`
	IpPosition  string `gorm:"column:ip_position;type:varchar(50);comment:ip属地" json:"ip_position" form:"ip_position"`
	gorm.Model
}

func (m *User) TableName() string {
	return "user"
}
