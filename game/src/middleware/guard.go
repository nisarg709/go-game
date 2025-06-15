package middleware

import (
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"eman/passport/game/src/api"
)

func Guard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token, ok := c.Get("guard").(*jwt.Token)

		if !ok {
			return echo.NewHTTPError(400, "Unauthorized")
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		c.Set("user", api.User{Id: claims["sub"].(string)})

		//@TODO: check for revoked tokens

		return next(c)
	}
}
