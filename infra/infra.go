package infra

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wzh/config"
)

var DB *gorm.DB
var RedisDb *redis.Client

func InitMysqlDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", config.Config.MysqlConfig.Username, config.Config.MysqlConfig.Password, config.Config.MysqlConfig.Host, config.Config.MysqlConfig.Port, config.Config.MysqlConfig.DbName, config.Config.MysqlConfig.Timeout)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func InitRedis() error {
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

func Init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := InitMysqlDB(); err != nil {
		panic(err)
	}

	if err := InitRedis(); err != nil {
		panic(err)
	}
}
