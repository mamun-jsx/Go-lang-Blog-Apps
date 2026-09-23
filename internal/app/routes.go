package app

import (
	"github.com/labstack/echo/v5"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/config"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/handlers"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/repositories"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/services"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {

	// repositories
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)
	commentRepo := repositories.NewCommentRepository(db)

	// services
	userSvc := services.NewUserService(userRepo)
	postSvc := services.NewPostService(postRepo, db)
	commentSvc := services.NewCommentService(commentRepo, db)

	// handlers
	authH := handlers.NewAuthHandler(userSvc)
	postH := handlers.NewPostHandler(postSvc)
	commentH := handlers.NewCommentHandler(commentSvc)

	// --- public routes ---
	api := e.Group("/api/v1")
	api.POST("/signup", authH.Signup)
	api.POST("/login", authH.Login)

	// --- protected routes (require JWT) ---
	protected := api.Group("", JWTMiddleware(cfg))

	// posts CRUD
	protected.POST("/posts", postH.CreatePost)
	protected.GET("/posts", postH.ListPosts)
	protected.GET("/posts/:id", postH.GetPost)
	protected.PUT("/posts/:id", postH.UpdatePost)
	protected.DELETE("/posts/:id", postH.DeletePost)

	// comments CRUD
	protected.POST("/posts/:id/comments", commentH.Add)
	protected.GET("/posts/:id/comments", commentH.List)
	protected.PUT("/comments/:id", commentH.Update)
	protected.DELETE("/comments/:id", commentH.Delete)
}
