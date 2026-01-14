package handlers

import (
	"homework04/models"
	"homework04/services"
	"homework04/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

// CreateComment 创建评论
func (h *CommentHandler) CreateComment(c *gin.Context) {
	// 获取当前用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// 绑定请求参数
	var req models.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, parseValidationErrors(err))
		return
	}

	// 创建评论
	comment, err := h.commentService.CreateComment(userID.(uint), req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, models.CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		UserID:    comment.UserID,
		Username:  "", // 需要从 User 关联获取
		PostID:    comment.PostID,
		CreatedAt: comment.CreatedAt,
	})
}

// GetCommentsByPostID 获取某篇文章的所有评论
func (h *CommentHandler) GetCommentsByPostID(c *gin.Context) {
	// 获取文章 ID
	idStr := c.Param("id")
	postID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid post id")
		return
	}

	// 获取评论列表
	comments, err := h.commentService.GetCommentsByPostID(uint(postID))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, comments)
}

// DeleteComment 删除评论
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	// 获取当前用户 ID
	currentUserID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// 获取要删除的评论 ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid comment id")
		return
	}

	// 删除评论（带权限验证）
	if err := h.commentService.DeleteCommentWithAuth(currentUserID.(uint), uint(id)); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"message": "Comment deleted successfully",
	})
}