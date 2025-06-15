package app

import (
	"net/http"
	"encoding/json"
)

type Responder struct {
}

func (responder *Responder) InternalServerError(w http.ResponseWriter, r *http.Request) {

	error := Error{
		Code:    500,
		Message: "Internal Server Error",
	}
	response := ErrorResponse{
		Error: error,
	}

	res, err := json.Marshal(response);
	if err != nil {
		res = []byte("")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(res)
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}
