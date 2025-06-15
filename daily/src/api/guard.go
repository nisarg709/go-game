package api

import (
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func Guard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token, ok := c.Get("guard").(*jwt.Token)

		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		c.Set("caller", &Caller{Id: claims["sub"].(string)})

		//@TODO: additional user & token validation

		return next(c)
	}
}
