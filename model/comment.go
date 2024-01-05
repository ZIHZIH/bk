package model

import (
	"gorm.io/gorm"
)

// 评论表
type Comment struct {
	Id            int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	ArticleId     int    `gorm:"column:article_id;type:int(11);comment:文章ID;NOT NULL" json:"article_id" form:"article_id"`
	CommentatorId int    `gorm:"column:commentator_id;type:int(11);comment:评论者ID;NOT NULL" json:"commentator_id" form:"commentator_id"`
	Content       string `gorm:"column:content;type:varchar(1024);comment:评论内容" json:"content" form:"content"`
	gorm.Model
}

func (m *Comment) TableName() string {
	return "comment"
}
