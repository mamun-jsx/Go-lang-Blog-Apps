package services

import (
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/repositories"
	"gorm.io/gorm"
)

type CommentService struct {
	repo *repositories.CommentRepository
	db   *gorm.DB
}

func NewCommentService(r *repositories.CommentRepository, db *gorm.DB) *CommentService {
	return &CommentService{repo: r, db: db}
}

