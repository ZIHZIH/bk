package config

import "github.com/spf13/viper"

var Config *ProjectConfig

type ProjectConfig struct {
	MysqlConfig *MysqlConfig `mapstructure:"mysql"`
	RedisConfig *RedisConfig `mapstructure:"cache"`
}

type MysqlConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"db_name"`
	Timeout  string `mapstructure:"timeout"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func readConfig() error {
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}
	return nil
}

func InitConfig() error {
	Config = new(ProjectConfig)
	Config.MysqlConfig = &MysqlConfig{}
	Config.RedisConfig = &RedisConfig{}
	err := readConfig()
	if err != nil {
		return err
	}
	return nil
}
