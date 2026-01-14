package handlers

import (
	"homework04/models"
	"homework04/services"
	"homework04/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

// CreatePost 创建文章
func (h *PostHandler) CreatePost(c *gin.Context) {
	// 获取当前用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// 绑定请求参数
	var req models.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, parseValidationErrors(err))
		return
	}

	// 设置用户 ID
	req.UserID = userID.(uint)

	// 创建文章
	post, err := h.postService.CreatePost(req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, models.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		Username:  "", // 需要从 User 关联获取
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	})
}

// GetPosts 获取所有文章列表
func (h *PostHandler) GetPosts(c *gin.Context) {
	// 获取 user_id 参数（可选）
	userIDStr := c.Query("user_id")
	var userID uint
	if userIDStr != "" {
		id, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			utils.Error(c, http.StatusBadRequest, "Invalid user_id")
			return
		}
		userID = uint(id)
	}

	// 获取文章列表
	posts, err := h.postService.GetPostsListByID(userID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, posts)
}

// GetPostByID 获取单个文章详情
func (h *PostHandler) GetPostByID(c *gin.Context) {
	// 获取文章 ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid post id")
		return
	}

	// 获取文章详情
	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 构造响应
	utils.Success(c, models.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		Username:  post.User.Username,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	})
}

// GetPostByTitle 根据标题获取文章详情（包含评论）
func (h *PostHandler) GetPostByTitle(c *gin.Context) {
	// 获取文章标题
	title := c.Param("title")
	if title == "" {
		utils.Error(c, http.StatusBadRequest, "Title is required")
		return
	}

	// 获取文章详情（包含评论）
	post, err := h.postService.GetPostDetailByTitle(title)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, post)
}

// UpdatePost 更新文章
func (h *PostHandler) UpdatePost(c *gin.Context) {
	// 获取当前用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// 获取文章 ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid post id")
		return
	}

	// 获取文章信息并验证权限
	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 验证是否为文章作者
	if post.UserID != userID.(uint) {
		utils.Error(c, http.StatusForbidden, "You can only update your own posts")
		return
	}

	// 绑定请求参数
	var req models.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(c, parseValidationErrors(err))
		return
	}

	// 更新文章
	updatedPost, err := h.postService.UpdatePost(uint(id), req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, models.PostResponse{
		ID:        updatedPost.ID,
		Title:     updatedPost.Title,
		Content:   updatedPost.Content,
		UserID:    updatedPost.UserID,
		Username:  updatedPost.User.Username,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: updatedPost.UpdatedAt,
	})
}

// DeletePost 删除文章
func (h *PostHandler) DeletePost(c *gin.Context) {
	// 获取当前用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// 获取文章 ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid post id")
		return
	}

	// 获取文章信息并验证权限
	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 验证是否为文章作者
	if post.UserID != userID.(uint) {
		utils.Error(c, http.StatusForbidden, "You can only delete your own posts")
		return
	}

	// 删除文章
	if err := h.postService.DeletePost(uint(id)); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, gin.H{
		"message": "Post deleted successfully",
	})
}
