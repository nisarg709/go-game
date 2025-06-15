package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type configuration struct {
	Port              string
	MongoDBUrl        string
	MongoDBName       string
	PublicKeyFilePath string
}

func config() *configuration {

	loadCongifuration()

	c := &configuration{
		viper.GetString("server.port"),
		viper.GetString("db.url"),
		viper.GetString("db.name"),
		viper.GetString("key.path"),
	}

	return c
}

func loadCongifuration() {

	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to load configuration: %v \n", err))
	}
}
