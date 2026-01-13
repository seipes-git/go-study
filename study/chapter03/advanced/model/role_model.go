package model

import (
	"math/rand"
	"time"
)

type Role struct {
	ID          uint
	Name        string `gorm:"uniqueIndex"` // Role name must be unique (e.g., "admin", "user", "editor")
	Description string
	Users       []User `gorm:"many2many:user_roles;"` // Many to Many: Role belongs to many users
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func GenerateRoles() []Role {
	roles := []Role{
		{Name: "admin", Description: "系统管理员"},
		{Name: "user", Description: "普通用户"},
		{Name: "editor", Description: "内容编辑"},
		{Name: "vip", Description: "VIP用户"},
		{Name: "guest", Description: "游客"},
	}
	now := time.Now()
	for i := range roles {
		roles[i].CreatedAt = now
		roles[i].UpdatedAt = now
	}
	return roles
}
