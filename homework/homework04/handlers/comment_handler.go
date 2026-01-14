package handlers

import (
	"github.com/gin-gonic/gin"
)

type CommentHandler struct{}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
}

// CreateComment 创建评论
func (h *CommentHandler) CreateComment(c *gin.Context) {
	// TODO: 实现创建评论逻辑
}

// GetCommentsByPostID 获取某篇文章的所有评论
func (h *CommentHandler) GetCommentsByPostID(c *gin.Context) {
	// TODO: 实现获取文章评论列表逻辑
}