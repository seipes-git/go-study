package main

import (
	"fmt"

	"gorm.io/gorm"

	"local/go_study/study/chapter03/advanced/db"
	"local/go_study/study/chapter03/advanced/model"
	"local/go_study/study/chapter03/advanced/service"
)

var Global_DB *gorm.DB

func main() {

	Global_DB, err := db.DBInit()
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
		return
	}

	// 迁移表结构
	err = Global_DB.AutoMigrate(&model.User{}, &model.Role{}, &model.Profile{}, &model.Product{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		fmt.Printf("表结构迁移失败: %v", err)
	} else {
		fmt.Println("数据库连接和表迁移成功！")
	}

	if err := service.DeleteData(Global_DB); err != nil {
		fmt.Printf("删除数据失败: %v", err)
	} else {
		fmt.Println("数据删除成功！")
	}

	if err := service.InsertData(Global_DB); err != nil {
		fmt.Printf("插入数据失败: %v", err)
	} else {
		fmt.Println("数据插入成功！")
	}
}
