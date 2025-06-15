package api

type Status struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type Data struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	HttpCode int         `json:"-"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Meta     interface{} `json:"meta,omitempty"`
}

func NewError(code int, message string) *Error {
	return &Error{
		HttpCode: 400,
		Code: code,
		Message: message,
	}
}
func (e Error) Error() string {
	return e.Message
}


func NewInternalServerError() ErrorResponse {
	return ErrorResponse{
		Error{
			Code:    10500,
			Message: "Internal Server Error",
		},
	}
}

func NewCustomError(message string, code int) ErrorResponse {

	if message == "" {
		message = "Bad Request"
	}

	if code == 0 {
		code = 10400
	}

	return ErrorResponse{
		Error{
			Code:    code,
			Message: message,
		},
	}
}

func NewBadRequest() ErrorResponse {
	return ErrorResponse{
		Error{
			Code:    10400,
			Message: "Bad Request",
		},
	}
}

func NewUnauthorized() ErrorResponse {
	return ErrorResponse{
		Error{
			Code:    10401,
			Message: "Unauthorized",
		},
	}
}

func NewForbidden() ErrorResponse {
	return ErrorResponse{
		Error{
			Code:    10403,
			Message: "Forbidden",
		},
	}
}

func NewNotFoundError() ErrorResponse {
	return ErrorResponse{
		Error{
			Code:    10404,
			Message: "Not Found",
		},
	}
}

func NewValidationError() ErrorResponse {
	return ErrorResponse{
		Error{
			Code:    11111,
			Message: "Validation Error",
		},
	}
}
