package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostRequest struct {
	UserID  uint   `json:"user_id" gorm:"not null;index"`
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"omitempty,min=1,max=200"`
	Content string `json:"content" binding:"omitempty,min=1"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostDetailResponse struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	UserID    uint              `json:"user_id"`
	Username  string            `json:"username"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Comments  []CommentResponse `json:"comments,omitempty"`
}
