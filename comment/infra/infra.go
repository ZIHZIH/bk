package infra

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
	"wzh/comment/config"
)

var MysqlDB *gorm.DB
var RedisDb *redis.Client
var Logger *log.Logger
var Mongodb *mongo.Client

func initMysqlDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", config.Config.MysqlConfig.Username, config.Config.MysqlConfig.Password, config.Config.MysqlConfig.Host, config.Config.MysqlConfig.Port, config.Config.MysqlConfig.DbName, config.Config.MysqlConfig.Timeout)
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func initRedis() error {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisConfig.Addr,     // redis地址
		Password: config.Config.RedisConfig.Password, // redis密码，没有则留空
		DB:       config.Config.RedisConfig.DB,       // 默认数据库，默认是0
	})

	_, err := RedisDb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func initLogger() error {
	file, err := os.OpenFile("./infra/wzh.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		return err
	}
	Logger = log.New(file, "wzh's bk:", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func initMongodb() error {
	// 设置客户端连接配置
	url := fmt.Sprintf("mongodb://%s:%d", config.Config.MongodbConfig.Host, config.Config.MongodbConfig.Port)
	clientOptions := options.Client().ApplyURI(url)

	// 连接到MongoDB
	var err error
	Mongodb, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// ping测试连接是否成功
	err = Mongodb.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func Init() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := initLogger(); err != nil {
		panic(err)
	}

	if err := initMysqlDB(); err != nil {
		panic(err)
	}

	if err := initRedis(); err != nil {
		panic(err)
	}

	if err := initMongodb(); err != nil {
		panic(err)
	}
}
