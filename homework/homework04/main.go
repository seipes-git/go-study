package main

import (
	"homework04/db"
	"homework04/models"
	"homework04/services"
	"log"
)

func main() {
	database, err := db.DBInit()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	log.Println("数据库连接成功")

	err = database.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	log.Println("数据库表迁移成功")

	err = services.SeedData(database)
	if err != nil {
		log.Fatalf("种子数据导入失败: %v", err)
	}

	log.Println("种子数据导入完成")
}
