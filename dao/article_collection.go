package dao

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"wzh/pkg/cache"
)

type ArticleMongodbDao struct {
	ArticleCollection *mongo.Collection
	Cache             *cache.Cache
}

func (articleMongodbDao *ArticleMongodbDao) ArticleInsertOne(ctx context.Context, data interface{}) (string, error) {
	res, err := articleMongodbDao.ArticleCollection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (articleMongodbDao *ArticleMongodbDao) ArticleFindOne(ctx context.Context, id string) (string, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	res := articleMongodbDao.ArticleCollection.FindOne(ctx, bson.M{"_id": objid})
	raw, err := res.Raw()
	if err != nil {
		return "", err
	}

	// 解码bson.Raw到map
	var doc map[string]interface{}
	if err = bson.Unmarshal(raw, &doc); err != nil {
		return "", err
	}

	// 将map编码为JSON字符串
	jsonBytes, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (articleMongodbDao *ArticleMongodbDao) ArticleDeleteOne(ctx context.Context, id string) error {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = articleMongodbDao.ArticleCollection.DeleteOne(ctx, bson.M{"_id": objid})
	if err != nil {
		return err
	}

	return err
}

func (articleMongodbDao *ArticleMongodbDao) ArticleListFind(ctx context.Context) ([]string, error) {
	var result []string

	cur, err := articleMongodbDao.ArticleCollection.Find(ctx, bson.M{})
	if err != nil {
		return []string{}, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var document bson.M
		if err = cur.Decode(&document); err != nil {
			return []string{}, err
		}

		// 将map编码为JSON字符串
		jsonBytes, err := json.Marshal(document)
		if err != nil {
			return []string{}, err
		}
		result = append(result, string(jsonBytes))
	}
	// 检查遍历过程中是否发生错误
	if err = cur.Err(); err != nil {
		return []string{}, err
	}

	return result, nil
}

func (articleMongodbDao *ArticleMongodbDao) ArticleUpdateOne(ctx context.Context, jsonStr string) (int64, error) {
	var updateData bson.M
	err := json.Unmarshal([]byte(jsonStr), &updateData)
	if err != nil {
		return -1, err
	}

	_, err = ArticleMongodbD.ArticleFindOne(ctx, updateData["_id"].(string))
	if err != nil {
		return -1, err
	}

	objid, err := primitive.ObjectIDFromHex(updateData["_id"].(string))
	if err != nil {
		return -1, err
	}
	delete(updateData, "_id")

	updateResult, err := articleMongodbDao.ArticleCollection.UpdateOne(ctx, bson.M{"_id": objid}, bson.M{"$set": updateData})
	if err != nil {
		return -1, err
	}

	return updateResult.ModifiedCount, nil
}

func NewArticleMongodbDao(ArticleCollection *mongo.Collection, Cache *cache.Cache) *ArticleMongodbDao {
	return &ArticleMongodbDao{
		ArticleCollection: ArticleCollection,
		Cache:             Cache,
	}
}
