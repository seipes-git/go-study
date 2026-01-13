package service

import (
	"fmt"
	"homework03/model"

	"gorm.io/gorm"
)

// DataService 数据服务，负责数据初始化和测试数据
type DataService struct {
	db             *gorm.DB
	userService    *UserService
	postService    *PostService
	commentService *CommentService
}

// NewDataService 创建数据服务
func NewDataService(db *gorm.DB) *DataService {
	return &DataService{
		db:             db,
		userService:    NewUserService(db),
		postService:    NewPostService(db),
		commentService: NewCommentService(db),
	}
}

// InitDatabase 初始化数据库表
func (s *DataService) InitDatabase() error {
	return s.db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}

// CreateTestData 创建测试数据
func (s *DataService) CreateTestData() error {
	// 清空数据库并重置自增 ID
	fmt.Println("清空数据库...")
	s.db.Exec("DELETE FROM blog_comments")
	s.db.Exec("DELETE FROM blog_posts")
	s.db.Exec("DELETE FROM blog_users")
	s.db.Exec("DELETE FROM sqlite_sequence WHERE name IN ('blog_comments', 'blog_posts', 'blog_users')")

	// 创建10个用户
	users := []model.User{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
		{Username: "user3", Email: "user3@example.com"},
		{Username: "user4", Email: "user4@example.com"},
		{Username: "user5", Email: "user5@example.com"},
		{Username: "user6", Email: "user6@example.com"},
		{Username: "user7", Email: "user7@example.com"},
		{Username: "user8", Email: "user8@example.com"},
		{Username: "user9", Email: "user9@example.com"},
		{Username: "user10", Email: "user10@example.com"},
	}

	if err := s.userService.CreateUsers(users); err != nil {
		return err
	}

	// 创建10篇文章，使用实际的用户 ID，不同用户有不等的文章数量
	posts := []model.Post{
		{Title: "Go语言入门", Content: "Go语言是一门简单高效的编程语言", UserID: users[0].ID},
		{Title: "Go并发编程", Content: "Go语言的并发模型非常优雅", UserID: users[0].ID},
		{Title: "Go标准库详解", Content: "Go标准库提供了丰富的功能", UserID: users[0].ID},
		{Title: "Gorm基础教程", Content: "Gorm是Go语言最流行的ORM框架", UserID: users[1].ID},
		{Title: "Gorm高级特性", Content: "Gorm的钩子函数和事务处理", UserID: users[1].ID},
		{Title: "微服务架构设计", Content: "微服务架构的优势与挑战", UserID: users[2].ID},
		{Title: "Docker容器技术", Content: "Docker让应用部署变得简单", UserID: users[2].ID},
		{Title: "MySQL性能调优", Content: "MySQL数据库性能优化的技巧", UserID: users[3].ID},
		{Title: "Vue.js实战", Content: "Vue.js前端开发实践", UserID: users[4].ID},
		{Title: "Kubernetes入门", Content: "Kubernetes是容器编排的事实标准", UserID: users[5].ID},
	}

	createdPosts, err := s.postService.CreatePosts(posts)
	if err != nil {
		return err
	}

	// 创建10条评论，使用实际的文章 ID
	comments := []model.Comment{
		{Content: "这篇文章写得很好！", PostID: createdPosts[0].ID},
		{Content: "学到了很多，感谢分享", PostID: createdPosts[0].ID},
		{Content: "Gorm确实很好用", PostID: createdPosts[1].ID},
		{Content: "微服务架构值得深入研究", PostID: createdPosts[2].ID},
		{Content: "Docker简化了部署流程", PostID: createdPosts[3].ID},
		{Content: "Kubernetes学习曲线有点陡", PostID: createdPosts[4].ID},
		{Content: "Redis性能确实很棒", PostID: createdPosts[5].ID},
		{Content: "MySQL优化很有用", PostID: createdPosts[6].ID},
		{Content: "Vue.js响应式系统很优雅", PostID: createdPosts[7].ID},
		{Content: "Spring Boot开箱即用", PostID: createdPosts[8].ID},
	}

	return s.commentService.CreateComments(comments)
}
