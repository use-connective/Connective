package server

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func Success(ctx *gin.Context, data any, message string, statusCode int) {
	ctx.JSON(statusCode, APIResponse{
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, err string, statusCode int) {
	ctx.JSON(statusCode, APIResponse{
		Error: err,
	})
}

func BadRequest(ctx *gin.Context, err string) {
	Error(ctx, err, 400)
}

func NotFound(ctx *gin.Context) {
	Error(ctx, "404 Not Found", 404)
}

func InternalError(ctx *gin.Context, err string) {
	Error(ctx, err, 500)
}

func Unauthorized(ctx *gin.Context) {
	Error(ctx, "User is unauthorized. Please login.", 401)
}
