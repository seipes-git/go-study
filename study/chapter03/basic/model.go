package basic

import (
	"gorm.io/gorm"
)

// type Model struct {
// 	UUID uint      `gorm:"primary_key"`
// 	Time time.Time `gorm:"column:my_time"`
// }

type TestUser struct {
	gorm.Model
	Name string `gorm:"default:seipes"`
	Age  uint8  `gorm:"comment:年龄"`
}

func TestUserCreate() {
	Global_DB.AutoMigrate(&TestUser{})
}
