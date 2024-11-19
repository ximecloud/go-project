package config

import (
	"fmt"
	"github.com/spf13/viper"
	"project/internal/adapter/storage/pg"
	"project/internal/adapter/storage/rds"
)

func Initial() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("cmd/config.yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetAppAddr() string {
	return fmt.Sprintf(":%v", viper.GetString("app.port"))
}
func GetPostgres() *pg.Config {
	return &pg.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}
}

func GetRedis() *rds.Config {
	conf := rds.Config{
		Host:     viper.GetString("cache.host"),
		Port:     viper.GetString("cache.port"),
		Password: viper.GetString("cache.password"),
		Database: viper.GetInt("cache.database"),
	}
	return &conf
}

func GetLogstash() string {
	return fmt.Sprintf("%s:%s", viper.GetString("logstash.host"), viper.GetString("logstash.port"))
}
