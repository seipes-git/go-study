package service

import (
	"fmt"
	"gorm.io/gorm"
	"homework03/model"
)

// CommentService 评论服务
type CommentService struct {
	db *gorm.DB
}

// NewCommentService 创建评论服务
func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

// CreateComment 创建评论
func (s *CommentService) CreateComment(comment *model.Comment) error {
	return s.db.Create(comment).Error
}

// CreateComments 批量创建评论
func (s *CommentService) CreateComments(comments []model.Comment) error {
	for i := range comments {
		if err := s.CreateComment(&comments[i]); err != nil {
			return fmt.Errorf("创建评论失败: %w", err)
		}
		fmt.Printf("创建评论: %s (ID: %d)\n", comments[i].Content, comments[i].ID)
	}
	return nil
}

// GetCommentByID 根据ID获取评论
func (s *CommentService) GetCommentByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	err := s.db.First(&comment, id).Error
	return &comment, err
}

// GetCommentsByPostID 获取某篇文章的所有评论
func (s *CommentService) GetCommentsByPostID(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := s.db.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}