package model

import "gorm.io/gorm"

// 点赞表
type Like struct {
	Id        int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	ArticleId int `gorm:"column:article_id;type:int(11);comment:文章ID;NOT NULL" json:"article_id" form:"article_id"`
	LikerId   int `gorm:"column:liker_id;type:int(11);comment:点赞者ID;NOT NULL" json:"liker_id" form:"liker_id"`
	gorm.Model
}

func (m *Like) TableName() string {
	return "like"
}
