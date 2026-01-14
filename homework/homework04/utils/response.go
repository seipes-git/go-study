package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Error:   message,
	})
}

func ValidationError(c *gin.Context, errors map[string]string) {
	c.JSON(422, Response{
		Code:    422,
		Message: "validation failed",
		Error:   errors,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}
