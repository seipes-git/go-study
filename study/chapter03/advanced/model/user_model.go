package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true"`
	Name      string
	Email     string
	Profile   Profile // Has One: One user has one profile
	Orders    []Order // Has Many: One user has many orders
	Roles     []Role  `gorm:"many2many:user_roles;"` // Many to Many: User has many roles through user_roles join table
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateUsers() []User {
	users := make([]User, 10)
	now := time.Now()
	for i := 0; i < 10; i++ {
		users[i] = User{
			Name:      fmt.Sprintf("用户%d", i+1),
			Email:     fmt.Sprintf("user%d@example.com", i+1),
			CreatedAt: now,
			UpdatedAt: now,
		}
	}
	return users
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 创建用户时，生成默认的昵称和手机号码
	if u.Profile.Nickname == "" { // 如果昵称为空，则生成一个默认的昵称
		u.Profile.Nickname = fmt.Sprintf("用户%d", u.ID)
	}

	if u.Profile.Phone == "" { // 如果手机号为空，则生成一个默认的手机号
		u.Profile.Phone = fmt.Sprintf("138%08d", u.ID)
	}

	return nil
}
