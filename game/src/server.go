package main

import (
	"eman/passport/game/src/container"
	"eman/passport/game/src/controllers"
	"eman/passport/game/src/errors"
	"eman/passport/game/src/middleware"
	"eman/passport/game/src/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"io/ioutil"
)

func main() {

	c := config()

	// Echo Init & Config
	server := echo.New()
	server.Debug = false
	server.HideBanner = true
	server.HTTPErrorHandler = errors.CustomHTTPErrorHandler

	// Middleware
	server.Use(mw.Recover())
	server.Use(mw.Logger())

	pem, err := ioutil.ReadFile(c.PublicKeyFilePath)

	if err != nil {
		panic(err);
	}
	key, _ := jwt.ParseRSAPublicKeyFromPEM(pem)

	config := mw.JWTConfig{
		ContextKey:    "guard",
		SigningKey:    key,
		SigningMethod: "RS256",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
	}

	jwt := mw.JWTWithConfig(config)
	revoked := middleware.Guard

	// DB
	session, err := mgo.Dial(c.MongoDBUrl)
	if err != nil {
		panic(err)
	}
	mongo := &repositories.MongoRepository{session}

	handler := controllers.Handler{
		&container.Container{
			PlaysRepository:             &repositories.MongoPlaysRepository{mongo},
			GamesRepository:             &repositories.MongoGamesRepository{mongo},
			UsersRepository:             &repositories.MongoUsersRepository{mongo},
			MilestonesRepository:        &repositories.MilestonesRepository{mongo},
			QuizRepository:              &repositories.QuizRepository{mongo},
			OddOneOutQuestionRepository: &repositories.OddOneOutQuestionRepository{mongo},
			OddOneOutImageRepository:    &repositories.OddOneOutImageRepository{mongo},
		},
	}

	// Routing
	g := server.Group("/v1/games", jwt, revoked)

	g.GET("/:game/:diff/start", handler.Start)
	g.GET("/:game/:diff/force", handler.Force)
	g.POST("/:id/complete", handler.Complete)
	g.GET("/:id/help", handler.Help)
	g.POST("/:id/help", handler.Help)
	g.GET("/:id/resume", handler.Resume)
	g.POST("/:id/resume", handler.Resume)
	g.GET("/:id/setup", handler.Setup)
	g.POST("/:id/check", handler.Check)

	g1 := server.Group("/v1.1/games", jwt, revoked)
	g1.GET("/:id/help", handler.Help)

	// Start
	server.Logger.Fatal(server.Start(c.Port))
}
