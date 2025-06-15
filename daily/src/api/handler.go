package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {

	code := http.StatusInternalServerError;
	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, NewError(code))
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
