package app

import (
	"crypto/rsa"
	"eman/passport/daily/src/controllers"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/go-redis/redis"
	"io/ioutil"
)

type App struct {
	Config  *Configuration
	Handler *controllers.Handler

	key   *rsa.PublicKey
	mongo *mgo.Session
	redis *redis.Client
}

func New(config *Configuration) *App {

	// Public key
	pem, err := ioutil.ReadFile(config.PublicKeyFilePath)
	check(err)

	key, err := jwt.ParseRSAPublicKeyFromPEM(pem)
	check(err)

	// Mongo DB
	mongo, err := mgo.Dial(config.MongoDBUrl)
	check(err)

	// Redis
	//redis := redis.NewClient(&redis.Options{
	//	Addr:     config.RedisAddr,
	//	Password: config.RedisPass,
	//	DB:       config.RedisDb,
	//})

	// Handler
	handler := controllers.New(mongo)

	return &App{
		Config:  config,
		Handler: handler,
		key:     key,
	}
}

func (app *App) GetPublicKey() *rsa.PublicKey {
	return app.key
}

func check(e error) {
	if e != nil {
		panic("app initialisation failed:" + e.Error());
	}
}
