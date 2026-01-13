package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"type:varchar(50);not null;uniqueIndex"`
	Email        string `gorm:"type:varchar(100);not null;uniqueIndex"`
	PostCount    int    `gorm:"default:0;comment:文章数量统计"`
	CreatedAt    int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Posts        []Post `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}