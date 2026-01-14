package services

import (
	"homework04/models"

	"gorm.io/gorm"
)

func SeedData(DB *gorm.DB) error {
	users := []models.User{
		{Username: "alice", Password: "password123", Email: "alice@example.com"},
		{Username: "bob", Password: "password123", Email: "bob@example.com"},
		{Username: "charlie", Password: "password123", Email: "charlie@example.com"},
		{Username: "david", Password: "password123", Email: "david@example.com"},
		{Username: "eve", Password: "password123", Email: "eve@example.com"},
		{Username: "frank", Password: "password123", Email: "frank@example.com"},
		{Username: "grace", Password: "password123", Email: "grace@example.com"},
		{Username: "henry", Password: "password123", Email: "henry@example.com"},
		{Username: "iris", Password: "password123", Email: "iris@example.com"},
		{Username: "jack", Password: "password123", Email: "jack@example.com"},
		{Username: "kate", Password: "password123", Email: "kate@example.com"},
		{Username: "leo", Password: "password123", Email: "leo@example.com"},
		{Username: "mary", Password: "password123", Email: "mary@example.com"},
		{Username: "nick", Password: "password123", Email: "nick@example.com"},
		{Username: "olivia", Password: "password123", Email: "olivia@example.com"},
		{Username: "peter", Password: "password123", Email: "peter@example.com"},
		{Username: "quinn", Password: "password123", Email: "quinn@example.com"},
		{Username: "rachel", Password: "password123", Email: "rachel@example.com"},
		{Username: "sam", Password: "password123", Email: "sam@example.com"},
		{Username: "tina", Password: "password123", Email: "tina@example.com"},
	}

	if err := DB.Create(&users).Error; err != nil {
		return err
	}

	posts := []models.Post{
		{Title: "First Post", Content: "This is my first post content", UserID: users[0].ID},
		{Title: "Hello World", Content: "Hello everyone, welcome to my blog", UserID: users[0].ID},
		{Title: "Go Programming", Content: "Go is a great programming language", UserID: users[1].ID},
		{Title: "Web Development", Content: "Web development is exciting", UserID: users[1].ID},
		{Title: "Database Design", Content: "Designing databases is important", UserID: users[2].ID},
		{Title: "API Design", Content: "RESTful API design principles", UserID: users[2].ID},
		{Title: "Microservices", Content: "Microservices architecture patterns", UserID: users[3].ID},
		{Title: "Cloud Computing", Content: "Cloud computing basics", UserID: users[3].ID},
		{Title: "DevOps", Content: "DevOps practices and tools", UserID: users[4].ID},
		{Title: "Containerization", Content: "Docker and Kubernetes", UserID: users[4].ID},
		{Title: "Security", Content: "Web security best practices", UserID: users[5].ID},
		{Title: "Testing", Content: "Unit testing strategies", UserID: users[5].ID},
		{Title: "CI/CD", Content: "Continuous integration and deployment", UserID: users[6].ID},
		{Title: "Frontend", Content: "Modern frontend frameworks", UserID: users[6].ID},
		{Title: "Backend", Content: "Backend development with Go", UserID: users[7].ID},
		{Title: "Data Structures", Content: "Essential data structures", UserID: users[7].ID},
		{Title: "Algorithms", Content: "Common algorithms", UserID: users[8].ID},
		{Title: "System Design", Content: "System design fundamentals", UserID: users[8].ID},
		{Title: "Performance", Content: "Performance optimization tips", UserID: users[9].ID},
		{Title: "Best Practices", Content: "Coding best practices", UserID: users[9].ID},
	}

	if err := DB.Create(&posts).Error; err != nil {
		return err
	}

	comments := []models.Comment{
		{Content: "Great post!", UserID: users[1].ID, PostID: posts[0].ID},
		{Content: "Very informative", UserID: users[2].ID, PostID: posts[0].ID},
		{Content: "Thanks for sharing", UserID: users[3].ID, PostID: posts[1].ID},
		{Content: "I learned a lot", UserID: users[4].ID, PostID: posts[2].ID},
		{Content: "Excellent explanation", UserID: users[5].ID, PostID: posts[3].ID},
		{Content: "This is helpful", UserID: users[6].ID, PostID: posts[4].ID},
		{Content: "Good job!", UserID: users[7].ID, PostID: posts[5].ID},
		{Content: "Interesting read", UserID: users[8].ID, PostID: posts[6].ID},
		{Content: "Well written", UserID: users[9].ID, PostID: posts[7].ID},
		{Content: "Clear and concise", UserID: users[0].ID, PostID: posts[8].ID},
		{Content: "Appreciate the effort", UserID: users[1].ID, PostID: posts[9].ID},
		{Content: "Very detailed", UserID: users[2].ID, PostID: posts[10].ID},
		{Content: "Useful information", UserID: users[3].ID, PostID: posts[11].ID},
		{Content: "Love this content", UserID: users[4].ID, PostID: posts[12].ID},
		{Content: "Keep it up!", UserID: users[5].ID, PostID: posts[13].ID},
		{Content: "Thanks for the tips", UserID: users[6].ID, PostID: posts[14].ID},
		{Content: "Amazing work", UserID: users[7].ID, PostID: posts[15].ID},
		{Content: "Helpful resource", UserID: users[8].ID, PostID: posts[16].ID},
		{Content: "Great insights", UserID: users[9].ID, PostID: posts[17].ID},
		{Content: "Looking forward to more", UserID: users[0].ID, PostID: posts[18].ID},
	}

	if err := DB.Create(&comments).Error; err != nil {
		return err
	}

	return nil
}
