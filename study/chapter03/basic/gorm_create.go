package basic

import "fmt"

func CreatedMultipleTest() {
	dbres := Global_DB.Create(&[]TestUser{
		{Name: "Alan", Age: 22},
		{Name: "Mike", Age: 23},
		{Name: "John", Age: 20},
		{Name: "Alice", Age: 18},
	})
	fmt.Println(dbres.Error, dbres.RowsAffected)
	if dbres.Error != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}
}

func CreatedSingleTest() {
	user := TestUser{Name: "Coco", Age: 15}
	dbres := Global_DB.Create(&user)
	if dbres.Error != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}

}
