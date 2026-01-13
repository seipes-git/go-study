package basic

import "fmt"

func DeleteTest() {
	var user TestUser
	// 软删除，会更新 deleted at
	// Global_DB.Model(&TestUser{}).Delete(&user, 10)
	// fmt.Println(errors.Is(dbres.Error, gorm.ErrRecordNotFound))

	// 软删除的数据可查看
	Global_DB.Unscoped().Find(&user, 1)
	fmt.Println(user)

	// 会直接删除数据
	// Global_DB.Unscoped().Delete(&user, 14)

}
