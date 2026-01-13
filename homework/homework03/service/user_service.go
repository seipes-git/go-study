package service

import (
	"fmt"
	"gorm.io/gorm"
	"homework03/model"
)

// UserService 用户服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建用户服务
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(user *model.User) error {
	return s.db.Create(user).Error
}

// CreateUsers 批量创建用户
func (s *UserService) CreateUsers(users []model.User) error {
	for i := range users {
		if err := s.CreateUser(&users[i]); err != nil {
			return fmt.Errorf("创建用户 %s 失败: %w", users[i].Username, err)
		}
		fmt.Printf("创建用户: %s (ID: %d)\n", users[i].Username, users[i].ID)
	}
	return nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := s.db.First(&user, id).Error
	return &user, err
}

// PrintUserWithPostsAndComments 获取并打印用户及其文章和评论信息
func (s *UserService) PrintUserWithPostsAndComments(userID uint) error {
	var user model.User
	err := s.db.Preload("Posts").Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		return err
	}

	fmt.Printf("用户: %s (ID: %d, 文章数: %d)\n", user.Username, user.ID, user.PostCount)
	for _, post := range user.Posts {
		fmt.Printf("  文章: %s (ID: %d)\n", post.Title, post.ID)
		for _, comment := range post.Comments {
			fmt.Printf("    评论: %s (ID: %d)\n", comment.Content, comment.ID)
		}
	}
	return nil
}