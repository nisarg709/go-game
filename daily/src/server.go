package main

import (
	"eman/passport/daily/src/api"
	"eman/passport/daily/src/app"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	cfg := app.Config()
	app := app.New(cfg)

	// Init
	e := echo.New()

	// Configuration
	e.Debug = false;
	e.HideBanner = true;
	e.HTTPErrorHandler = api.HTTPErrorHandler

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		ContextKey:    "guard",
		SigningKey:    app.GetPublicKey(),
		SigningMethod: "RS256",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
	}))
	e.Use(api.Guard)

	// Routes
	e.GET("/v1/launch", app.Handler.Daily)

	// Start
	e.Logger.Fatal(e.Start(cfg.Port))
}
