package model

import (
	"gorm.io/gorm"
)

// Post 文章模型
type Post struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"type:varchar(200);not null"`
	Content     string    `gorm:"type:text;not null"`
	UserID      uint      `gorm:"not null;index;comment:用户ID"`
	User        User      `gorm:"foreignKey:UserID"`
	CommentStatus string  `gorm:"type:varchar(20);default:'有评论';comment:评论状态"`
	CreatedAt   int64     `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64     `gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Comments    []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}

// PostWithCommentCount 带评论数量的文章（用于查询）
type PostWithCommentCount struct {
	Post
	CommentCount int64
}

// AfterCreate Hook: 创建文章后自动更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// AfterDelete Hook: 删除文章后自动减少用户的文章数量
func (p *Post) AfterDelete(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count - ?", 1)).Error
}