package main

import (
	"homework04/config"
	"homework04/db"
	"homework04/handlers"
	"homework04/middleware"
	"homework04/models"
	"homework04/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	config.LoadConfig()
	utils.InitLogger()

	// 初始化数据库
	database, err := db.DBInit()
	if err != nil {
		utils.LogErrorf("数据库连接失败: %v", err)
		log.Fatalf("数据库连接失败: %v", err)
	}

	utils.LogInfo("数据库连接成功")

	// 自动迁移模型
	err = database.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		utils.LogErrorf("数据库迁移失败: %v", err)
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化 Gin 路由
	r := gin.Default()

	// 初始化 handlers
	userHandler := handlers.NewUserHandler()
	postHandler := handlers.NewPostHandler()
	commentHandler := handlers.NewCommentHandler()

	// 公开路由（无需认证）
	public := r.Group("/api/v1")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
		public.GET("/posts", postHandler.GetPosts)
		public.GET("/posts/:id", postHandler.GetPostByID)
		public.GET("/posts/:id/comments", commentHandler.GetCommentsByPostID)
	}

	// 需要认证的路由
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/posts", postHandler.CreatePost)
		protected.PUT("/posts/:id", postHandler.UpdatePost)
		protected.DELETE("/posts/:id", postHandler.DeletePost)
		protected.POST("/posts/:id/comments", commentHandler.CreateComment)
	}

	// 启动服务器
	utils.LogInfo("服务器启动在 :8080")
	if err := r.Run(":8080"); err != nil {
		utils.LogErrorf("服务器启动失败: %v", err)
		log.Fatalf("服务器启动失败: %v", err)
	}
}
