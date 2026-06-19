package response

import "github.com/gin-gonic/gin"

type Body struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, Body{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, statusCode int, message string, errors interface{}) {
	ctx.JSON(statusCode, Body{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
