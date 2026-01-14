package services

import (
	"errors"
	"homework04/models"
	"homework04/utils"

	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

func (s *CommentService) CreateComment(userID uint, req models.CreateCommentRequest) (*models.Comment, error) {
	// 检查文章是否存在
	var post models.Post
	if err := s.db.First(&post, req.PostID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppError(404, "Post not found")
		}
		return nil, err
	}

	// 创建评论
	comment := models.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  req.PostID,
	}

	if err := s.db.Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (s *CommentService) GetCommentsByPostID(postID uint) ([]models.CommentResponse, error) {
	// 检查文章是否存在
	var post models.Post
	if err := s.db.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppError(404, "Post not found")
		}
		return nil, err
	}

	var comments []models.Comment

	// 查询评论并预加载用户信息
	if err := s.db.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		return nil, err
	}

	// 转换为 Response 结构
	responses := make([]models.CommentResponse, len(comments))
	for i, comment := range comments {
		responses[i] = models.CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    comment.UserID,
			Username:  comment.User.Username,
			PostID:    comment.PostID,
			CreatedAt: comment.CreatedAt,
		}
	}

	return responses, nil
}

func (s *CommentService) DeleteComment(id uint) error {
	comment, err := s.GetCommentByID(id)
	if err != nil {
		return err
	}

	if err := s.db.Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}

func (s *CommentService) DeleteCommentWithAuth(userID, commentID uint) error {
	comment, err := s.GetCommentByID(commentID)
	if err != nil {
		return err
	}

	// 验证权限：只能删除自己的评论
	if comment.UserID != userID {
		return utils.NewAppError(403, "You can only delete your own comments")
	}

	if err := s.db.Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}

func (s *CommentService) GetCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	if err := s.db.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppError(404, "Comment not found")
		}
		return nil, err
	}
	return &comment, nil
}