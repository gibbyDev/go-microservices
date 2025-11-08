package config

// import (
// 	"os"
// 	"strconv"
// )

// type Env struct {
// 	Port               string
// 	AuthServiceURL     string
// 	UserServiceURL     string
// 	PostServiceURL     string
// 	CommentServiceURL  string
// 	LikeServiceURL     string
// 	FollowServiceURL   string
// 	NotificationURL    string
// 	RateLimitPerMinute int
// }

// func LoadEnv() *Env {
// 	rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT_PER_MINUTE"))
// 	if err != nil {
// 		rateLimit = 60 // default value
// 	}
// 	return &Env{
// 		Port:               getEnv("PORT", "8080"),
// 		AuthServiceURL:     getEnv("AUTH_SERVICE_URL", "http://localhost:8001"),
// 		UserServiceURL:     getEnv("USER_SERVICE_URL", "http://localhost:8002"),
// 		PostServiceURL:     getEnv("POST_SERVICE_URL", "http://localhost:8003"),
// 		CommentServiceURL:  getEnv("COMMENT_SERVICE_URL", "http://localhost:8004"),
// 		LikeServiceURL:     getEnv("LIKE_SERVICE_URL", "http://localhost:8005"),
// 		FollowServiceURL:   getEnv("FOLLOW_SERVICE_URL", "http://localhost:8006"),
// 		NotificationURL:    getEnv("NOTIFICATION_SERVICE_URL", "http://localhost:8007"),
// 		RateLimitPerMinute: rateLimit,
// 	}
// }
