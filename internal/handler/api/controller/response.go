package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type JsonResponse[T any] struct {
	Data    T              `json:"data,omitempty"`
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Success bool           `json:"success,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func customErrorResponse(ctx *fiber.Ctx, err error) error {
	ctx.Status(http.StatusBadRequest)
	return ctx.JSON(JsonResponse[any]{
		Error: &ErrorResponse{
			ErrorMessage: err.Error(),
		},
	})
}
