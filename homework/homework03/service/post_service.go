package service

import (
	"fmt"
	"gorm.io/gorm"
	"homework03/model"
)

// PostService 文章服务
type PostService struct {
	db *gorm.DB
}

// NewPostService 创建文章服务
func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

// CreatePost 创建文章
func (s *PostService) CreatePost(post *model.Post) error {
	return s.db.Create(post).Error
}

// CreatePosts 批量创建文章
func (s *PostService) CreatePosts(posts []model.Post) ([]model.Post, error) {
	for i := range posts {
		if err := s.CreatePost(&posts[i]); err != nil {
			return nil, fmt.Errorf("创建文章 %s 失败: %w", posts[i].Title, err)
		}
		fmt.Printf("创建文章: %s (ID: %d)\n", posts[i].Title, posts[i].ID)
	}
	return posts, nil
}

// GetPostByID 根据ID获取文章
func (s *PostService) GetPostByID(id uint) (*model.Post, error) {
	var post model.Post
	err := s.db.First(&post, id).Error
	return &post, err
}

// PrintPostWithMostComments 获取并打印评论数量最多的文章
func (s *PostService) PrintPostWithMostComments() error {
	var post model.PostWithCommentCount
	err := s.db.Model(&model.Post{}).
		Select("blog_posts.*, count(blog_comments.id) as comment_count").
		Joins("LEFT JOIN blog_comments ON blog_comments.post_id = blog_posts.id").
		Group("blog_posts.id").
		Order("comment_count DESC").
		First(&post).Error
	if err != nil {
		return err
	}
	fmt.Printf("评论最多的文章: %s (ID: %d)\n", post.Title, post.ID)
	fmt.Printf("评论数量: %d\n", post.CommentCount)
	return nil
}