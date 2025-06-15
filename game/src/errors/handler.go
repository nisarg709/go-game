package errors

import (
	"eman/passport/game/src/api"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var response api.ErrorResponse
	var code = http.StatusInternalServerError
	var msg interface{}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)
		}
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}

	if e, ok := err.(*api.Error); ok {
		code = e.HttpCode

		error := api.Error{
			Code:    e.Code,
			Message: e.Message,
		}

		c.Logger().Error(error)

		response = api.ErrorResponse{
			Error: error,
		}
	} else if msg, ok := msg.(string); ok {
		response = api.ErrorResponse{
			Error: api.Error{
				Code:    code,
				Message: msg,
			},
		}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, response)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
