package infra

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wzh/config"
)

var DB *gorm.DB

func InitMysqlDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", config.Config.MysqlConfig.Username, config.Config.MysqlConfig.Password, config.Config.MysqlConfig.Host, config.Config.MysqlConfig.Port, config.Config.MysqlConfig.DbName, config.Config.MysqlConfig.Timeout)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Init() error {
	if err := config.InitConfig(); err != nil {
		return err
	}

	if err := InitMysqlDB(); err != nil {
		return err
	}

	return nil
}
