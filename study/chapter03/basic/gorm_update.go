package basic

// 先查到自己要的数据后再进行修改
func UpdateTest() {
	//	update		只更新你选择的字段
	//	updates		更新所有字段，此时有两种形式，一种为MAP，一种为结构体，结构体零值不参与更新
	//	save		无论如何都更新所有的内容，包括零值

	// Global_DB.Model(&TestUser{}).Where("name = ?", "Alan").Update("name", "seipes")

	// Global_DB.Model(&TestUser{}).Where("name = ?", "seipes").Updates(map[string]interface{}{"name": 0})

	Global_DB.Model(&TestUser{}).Where("name = ?", "0").Updates(TestUser{Name: "seipes", Age: 18})
}
