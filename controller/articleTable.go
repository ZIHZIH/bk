package controller

import (
	"errors"
	"time"
)

// 模拟自增id
var articleTotal = 0

// 模拟数据库存储（key：id）
var articleMap map[int]*ArticleRecord

type ArticleRecord struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Label      string    `json:"label"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
}

func init() {
	articleMap = make(map[int]*ArticleRecord)
}

// GetArticle 根据id查询文章记录
func GetArticle(id int) (*ArticleRecord, error) {
	if v, ok := articleMap[id]; ok {
		return v, nil
	}
	return nil, errors.New("article is not exist")
}

// CreatArticle 创建新的文章记录
func CreatArticle(a *ArticleRecord) (*ArticleRecord, error) {
	articleTotal++
	a.Id = articleTotal
	a.CreateTime = time.Now()
	articleMap[articleTotal] = a
	return a, nil
}

// DeleteArticle 根据文章id对文章记录进行删除
func DeleteArticle(id int) error {
	delete(articleMap, id)
	return nil
}

// UpdateArticle 对文章进行更新
func UpdateArticle(a *ArticleRecord) (*ArticleRecord, error) {
	articleMap[a.Id] = a
	return a, nil
}

// ListArticle 获取所有文章列表
func ListArticle() ([]*ArticleRecord, error) {
	result := make([]*ArticleRecord, 0)
	for _, v := range articleMap {
		result = append(result, v)
	}
	return result, nil
}
