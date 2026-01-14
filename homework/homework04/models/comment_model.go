package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	PostID    uint      `json:"post_id" gorm:"not null;index"`
	Post      Post      `json:"post,omitempty" gorm:"foreignKey:PostID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1,max=1000"`
	PostID  uint   `json:"post_id" binding:"required"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	PostID    uint      `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}
