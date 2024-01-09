package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
	"wzh/dao"
)

// ArticleGet 文章的获取
func ArticleGet(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dont have parameter：id"})
		return
	}

	//recordId, err := strconv.Atoi(articleId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	//result, err := dao.ArticleD.GetArticle(c, recordId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	result, err := dao.ArticleMongodbD.ArticleFindOne(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// ArticleUpdate 文章的更新
func ArticleUpdate(c *gin.Context) {
	//temp := new(model.Article)
	//err := c.ShouldBindJSON(temp)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	b := make([]byte, c.Request.ContentLength)
	_, err := c.Request.Body.Read(b)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("json", string(b))
	result, err := dao.ArticleMongodbD.ArticleUpdateOne(c, string(b))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ArticleDelete 文章的删除
func ArticleDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dont have parameter：id"})
		return
	}

	//recordId, err := strconv.Atoi(articleId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//err = dao.ArticleD.DeleteArticle(c, recordId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	err := dao.ArticleMongodbD.ArticleDeleteOne(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// ArticleCreat 文章的创建
func ArticleCreat(c *gin.Context) {
	//temp := new(model.Article)
	//err := c.ShouldBindJSON(temp)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//result, err := dao.ArticleD.CreatArticle(c, temp)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	b := make([]byte, c.Request.ContentLength)
	_, err := c.Request.Body.Read(b)
	if err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data bson.D
	err = bson.UnmarshalExtJSON(b, true, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := dao.ArticleMongodbD.ArticleInsertOne(c, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// ListArticle 列出所有文章
func ListArticle(c *gin.Context) {
	//resp, err := dao.ArticleD.ListArticle(c)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	result, err := dao.ArticleMongodbD.ArticleListFind(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
