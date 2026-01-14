package handlers

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Register 用户注册
func (h *UserHandler) Register(c *gin.Context) {
	// TODO: 实现用户注册逻辑
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	// TODO: 实现用户登录逻辑
}