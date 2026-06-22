package response

import "github.com/gin-gonic/gin"

type SuccessBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, SuccessBody{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, statusCode int, message string, errors interface{}) {
	ctx.JSON(statusCode, ErrorBody{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
