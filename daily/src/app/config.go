package app

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Port              string
	MongoDBUrl        string
	PublicKeyFilePath string
	RedisAddr         string
	RedisPass         string
	RedisDb           int
}

func Config() *Configuration {

	loadCongifuration()

	c := &Configuration{
		viper.GetString("server.port"),
		viper.GetString("db.url"),
		viper.GetString("key.path"),
		viper.GetString("redis.addr"),
		viper.GetString("redis.pass"),
		viper.GetInt("redis.db"),
	}

	return c
}

func loadCongifuration() {

	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic("app initialisation failed:" + err.Error());
	}
}
