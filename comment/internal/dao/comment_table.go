package dao

import (
	"bk/comment/infra"
	model "bk/comment/internal/dao/models"
	"context"
	"log"
)

type CommentDao struct {
	Logger *log.Logger
}

// CreatComment 创建点赞记录
func (c *CommentDao) CreatComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	if err := infra.MysqlDB.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentByArticleId 通过文章id获得该文章的评论信息
func (c *CommentDao) GetCommentByArticleId(ctx context.Context, articleId int) ([]*model.Comment, error) {
	var res []*model.Comment
	if err := infra.MysqlDB.Where("article_id = ?", articleId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func NewCommentDao(Logger *log.Logger) *CommentDao {
	return &CommentDao{Logger: Logger}
}
