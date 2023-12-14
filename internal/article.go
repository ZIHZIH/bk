package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"wzh/controller"
)

// ArticleCreate 文章的创建
func ArticleCreate(res http.ResponseWriter, req *http.Request) {
	fmt.Println("新建了一个文章")
	defer func() { _ = req.Body.Close() }()
	content, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	article := &controller.Article{}
	err = json.Unmarshal(content, article)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	// 将读取到到内容写入到数据库
	result, err := controller.CreatArticle(article)
	if err != nil {
		res.Write([]byte(err.Error()))
	}

	// 填写响应消息
	res.Write([]byte(result))
}

// ArticleDelete 文章的删除
func ArticleDelete(res http.ResponseWriter, req *http.Request) {
	fmt.Println("文章被删除")
	defer func() { _ = req.Body.Close() }()

	// 取出要删除到记录id
	data := req.URL.Query()
	id := data.Get("id")

	// 删除响应到文章记录
	isDelete, err := controller.DeleteArticle(id)
	if err != nil {
		res.Write([]byte(err.Error()))
	}

	if !isDelete {
		res.Write(httpResponseFailMessage)
	}
	res.Write(httpResponseSuccessMessage)
}

// ArticleUpdate 文章的更新
func ArticleUpdate(res http.ResponseWriter, req *http.Request) {
	fmt.Println("文章被更新")
	defer func() { _ = req.Body.Close() }()
	content, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	// 将响应的内容更新至数据库
	record := &controller.ArticleRecord{}
	json.Unmarshal(content, record)

	result, err := controller.UpdateArticle(record)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	// 将更新后的内容写入到消息主体
	res.Write([]byte(result))
}

// ArticleGet 文章的获取
func ArticleGet(res http.ResponseWriter, req *http.Request) {
	fmt.Println("文章被查询")
	defer func() { _ = req.Body.Close() }()
	data := req.URL.Query()
	id := data.Get("id")
	// 根据id去数据库进行响应的查询
	recordId, err := strconv.Atoi(id)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	result, err := controller.GetArticle(recordId)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	// 将响应的结果写入的响应消息中
	res.Write([]byte(result))
}

// ArticleProcess 对文章处理
func ArticleProcess(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// 文章获取
		ArticleGet(res, req)
	case "POST":
		// 文章的创建
		ArticleCreate(res, req)
	case "UPDATE", "PUT":
		// 文章的更新
		ArticleUpdate(res, req)
	case "DELETE":
		// 文章的删除
		ArticleDelete(res, req)
	}
}
