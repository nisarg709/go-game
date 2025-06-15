package service

import (
	"github.com/joho/godotenv"
	"github.com/caarlos0/env"
	"flag"
)

type Configuration struct {
	Port          int    `env:"PORT" envDefault:8080`
	LogExecution  bool   `env:"LOG_EXECUTION" envDefault:true`
	RedisEndpoint string `env:"REDIS_ENDPOINT"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDatabase int    `env:"REDIS_DATABASE" envDefault:0`
	MongoEndpoint string `env:"MONGO_ENDPOINT"`
	MongoDatabase string `env:"MONGO_DATABASE"`
}

func NewConfiguration() *Configuration {

	var path = flag.String("env", ".env", "Relative path to .env file if provided")
	flag.Parse()

	loadEnvironmentConfiguration(path)

	c := new(Configuration)
	err := env.Parse(c)

	if err != nil {
		panic(err)
	}

	//@TODO: validate if all the needed configuration options are set
	return c
}

func loadEnvironmentConfiguration(path *string) {

	var err error;

	if(path != nil) {
		err = godotenv.Load(*path)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		// We are only trying to load .env files to add additional variables
		// If the .env files does not exist all the needed variables MUST BE
		// exported and will be available with os.Getenv, IF NOT the error will
		// be handled when initializing the configuration struct
	}
}
