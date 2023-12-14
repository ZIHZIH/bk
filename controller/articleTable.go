package controller

import (
	"encoding/json"
	"errors"
	"strconv"
)

// 模拟自增id
var articleTotal = 0

// 模拟数据库存储（key：id）
var articleMap map[int]*Article

type Article struct {
	Userid string `json:"userid"`
	Text   string `json:"text"`
}

type ArticleRecord struct {
	Id      string `json:"id"`
	Article *Article
}

func init() {
	articleMap = make(map[int]*Article)
}

// GetArticle 根据id查询文章记录
func GetArticle(id int) (string, error) {
	if v, ok := articleMap[id]; ok {
		record := &ArticleRecord{
			Id:      strconv.Itoa(id),
			Article: v,
		}
		b, err := json.Marshal(record)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return "", errors.New("article is not exist")
}

// CreatArticle 创建新的文章记录
func CreatArticle(a *Article) (string, error) {
	articleTotal++
	articleMap[articleTotal] = a

	articleResp := ArticleRecord{
		Id:      strconv.Itoa(articleTotal),
		Article: articleMap[articleTotal],
	}

	b, err := json.Marshal(articleResp)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// DeleteArticle 根据文章id对文章记录进行删除
func DeleteArticle(id string) (bool, error) {
	key, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}

	delete(articleMap, key)
	return true, nil
}

// UpdateArticle 对文章进行更新
func UpdateArticle(record *ArticleRecord) (string, error) {
	key, err := strconv.Atoi(record.Id)
	if err != nil {
		return "", err
	}

	if _, ok := articleMap[key]; ok {
		articleMap[key] = record.Article
		b, err := json.Marshal(record)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	return "", errors.New("article is not exist")
}
