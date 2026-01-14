package services

import (
	"errors"
	"homework04/models"
	"homework04/utils"

	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) CreatePost(req models.CreatePostRequest) (*models.Post, error) {
	// 检查该用户是否已有相同标题的文章
	var existingPost models.Post
	if err := s.db.Where("user_id = ? AND title = ?", req.UserID, req.Title).First(&existingPost).Error; err == nil {
		return nil, utils.NewAppError(400, "Post with the same title already exists for this user")
	}

	// 创建文章
	post := models.Post{
		UserID:  req.UserID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := s.db.Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := s.db.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppError(404, "Post not found")
		}
		return nil, err
	}
	return &post, nil
}

func (s *PostService) GetPostsListByID(id uint) ([]models.PostResponse, error) {
	var posts []models.Post

	// 如果提供了 id，则查询特定用户的文章，否则查询所有文章
	query := s.db.Preload("User")
	if id > 0 {
		query = query.Where("user_id = ?", id)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	// 转换为 Response 结构
	responses := make([]models.PostResponse, len(posts))
	for i, post := range posts {
		responses[i] = models.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserID:    post.UserID,
			Username:  post.User.Username,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
	}

	return responses, nil
}

func (s *PostService) GetPostDetailByTitle(title string) (*models.PostDetailResponse, error) {
	var post models.PostDetailResponse
	if err := s.db.Where("title = ?", title).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppError(404, "Post not found")
		}
		return nil, err
	}
	return &post, nil
}

func (s *PostService) UpdatePost(id uint, req models.UpdatePostRequest) (*models.Post, error) {
	post, err := s.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	// 如果更新标题，检查标题是否已存在（排除当前文章）
	if req.Title != "" && req.Title != post.Title {
		var existingPost models.Post
		if err := s.db.Where("title = ? AND id != ?", req.Title, id).First(&existingPost).Error; err == nil {
			return nil, utils.NewAppError(400, "Post with the same title already exists")
		}
		post.Title = req.Title
	}

	// 更新内容
	if req.Content != "" {
		post.Content = req.Content
	}

	if err := s.db.Save(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) DeletePost(id uint) error {
	post, err := s.GetPostByID(id)
	if err != nil {
		return err
	}

	if err := s.db.Delete(&post).Error; err != nil {
		return err
	}

	return nil
}
