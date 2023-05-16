package dto_response

import "net/http"

type Response struct {
	Data interface{} `json:"data"`
} // @name Response

type SuccessResponse struct {
	Message string `json:"message" example:"OK"`
} // @name SuccessResponse

type Error struct {
	Domain  string `json:"domain"`
	Message string `json:"message"`
} // @name Error

type ErrorResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
} // @name ErrorResponse

type DataResponse map[string]interface{} // @name DataResponse

type PaginationResponse struct {
	Total int         `json:"total" example:"24"`
	Page  *int        `json:"page" example:"1"`
	Limit *int        `json:"limit" example:"10"`
	Nodes interface{} `json:"nodes"`
} // @name PaginationResponse

func NewBadRequestResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: message,
		Errors:  []Error{},
	}
}

func NewBadRequestResponseP(message string) *ErrorResponse {
	r := NewBadRequestResponse(message)
	return &r
}

func NewUnauthorizedResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: message,
		Errors:  []Error{},
	}
}

func NewUnauthorizedResponseP(message string) *ErrorResponse {
	r := NewUnauthorizedResponse(message)
	return &r
}

func NewForbiddenResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusForbidden,
		Message: message,
		Errors:  []Error{},
	}
}

func NewForbiddenResponseP(message string) *ErrorResponse {
	r := NewForbiddenResponse(message)
	return &r
}

func NewNotFoundResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusNotFound,
		Message: message,
		Errors:  []Error{},
	}
}

func NewNotFoundResponseP(message string) *ErrorResponse {
	r := NewNotFoundResponse(message)
	return &r
}

func NewInternalServerErrorResponse() ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Errors:  []Error{},
	}
}

func NewInternalServerErrorResponseP() *ErrorResponse {
	r := NewInternalServerErrorResponse()
	return &r
}
