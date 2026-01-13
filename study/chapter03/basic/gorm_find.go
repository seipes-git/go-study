package basic

import "fmt"

type UserDemo struct {
	Name string
	Age  int
}

func FindTest() {
	// result := map[string]interface{}{}
	// var User []TestUser
	var UserDemo []UserDemo

	// Global_DB.Model(&TestUser{}).First(&User)
	// Global_DB.Model(&TestUser{}).Take(&User)
	// Global_DB.Model(&TestUser{}).Find(&User)
	// Global_DB.Model(&TestUser{}).Find(&UserDemo)
	// dbres := Global_DB.Model(&TestUser{}).Last(&User, 15)
	// fmt.Println(errors.Is(dbres.Error, gorm.ErrRecordNotFound))

	// fmt.Println(result)
	// fmt.Println(User)

	// Global_DB.Where("name = ?", "John").First(&User)

	// fmt.Println(User)

	fmt.Println(UserDemo)
}
