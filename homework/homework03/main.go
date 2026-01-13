package main

import (
	"fmt"
	"log"

	"homework03/db"
	"homework03/service"
)

func main() {
	// 初始化数据库连接
	database, err := db.DBInit()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 初始化服务层
	dataService := service.NewDataService(database)
	userService := service.NewUserService(database)
	postService := service.NewPostService(database)

	// 题目1: 创建数据库表
	fmt.Println("=== 题目1: 创建数据库表 ===")
	if err := dataService.InitDatabase(); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	fmt.Println("数据库表创建成功!")

	// 创建测试数据
	fmt.Println("\n=== 创建测试数据 ===")
	if err := dataService.CreateTestData(); err != nil {
		log.Fatalf("创建测试数据失败: %v", err)
	}

	// 题目2: 关联查询
	fmt.Println("\n=== 题目2: 关联查询 ===")

	// 2.1 查询某个用户发布的所有文章及其对应的评论信息
	fmt.Println("2.1 查询用户发布的所有文章及其评论:")
	if err := userService.PrintFirstUserWithPostsAndComments(); err != nil {
		fmt.Printf("查询失败: %v\n", err)
	}

	// 2.2 查询评论数量最多的文章信息
	fmt.Println("\n2.2 查询评论数量最多的文章:")
	if err := postService.PrintPostWithMostComments(); err != nil {
		fmt.Printf("查询失败: %v\n", err)
	}
}
