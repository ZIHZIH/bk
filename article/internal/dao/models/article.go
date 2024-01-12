package model

import "gorm.io/gorm"

// 文章表
type Article struct {
	Id       int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	AuthorId int    `gorm:"column:author_id;type:int(11);comment:作者ID;NOT NULL" json:"author_id"`
	Title    string `gorm:"column:title;type:varchar(50);comment:文章标题;NOT NULL" json:"title"`
	Content  string `gorm:"column:content;type:varchar(1024);comment:文章内容;NOT NULL" json:"content"`
	Label    string `gorm:"column:label;type:varchar(50);comment:文章标签" json:"label"`
	Status   int    `gorm:"column:Status;type:tinyint(4);comment:状态;NOT NULL" json:"status"`
	gorm.Model
}

func (m *Article) TableName() string {
	return "article"
}
