package api

import "net/http"

type Status struct {
	Status bool `json:"status"`
}

type Data struct {
	Data interface{} `json:"data"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func NewError(code int) *ErrorResponse {

	var (
		id  int
		msg string
	)

	switch code {
	case
		http.StatusInternalServerError,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound:

		id = code
		msg = http.StatusText(code)
		break
	default:
		id = http.StatusInternalServerError
		msg = http.StatusText(code)
	}

	return &ErrorResponse{
		Error{
			Code:    id,
			Message: msg,
		},
	}
}
