package model

import (
	"fmt"
	"time"
)

type Profile struct {
	ID        uint
	UserID    uint `gorm:"uniqueIndex"` // Foreign key to user, unique to enforce one-to-one
	Nickname  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateProfiles(users []User) []Profile {
	profiles := make([]Profile, len(users))
	now := time.Now()
	for i, user := range users {
		profiles[i] = Profile{
			UserID:    user.ID, // 使用数据库生成的真实User.ID
			Nickname:  fmt.Sprintf("昵称%d", i+1),
			Phone:     fmt.Sprintf("138%08d", i+1), // 13800000001 ~ 13800000010
			Address:   fmt.Sprintf("地址%d号", i+1),
			CreatedAt: now,
			UpdatedAt: now,
		}
	}
	return profiles
}
