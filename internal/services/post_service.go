package services

import (
	"time"

	"github.com/google/uuid"

	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/models"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/repositories"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/utils"
	"gorm.io/gorm"
)

type PostService struct {
	repo *repositories.PostRepository
	db   *gorm.DB
}

func NewPostService(r *repositories.PostRepository, db *gorm.DB) *PostService {
	return &PostService{
		repo: r,
		db:   db,
	}
}

// create a post
func (s *PostService) Create(authorID uuid.UUID, title, content string) (*models.Post, error) {

	now := time.Now()
	post := &models.Post{
		ID:         uuid.New(),
		AuthorID:   authorID,
		Title:      title,
		Slug:       utils.MakeSlugSimple(title + "-" + uuid.NewString()[:6]),
		Status:     "published",
		Visibility: "public",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(post).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return post, nil
}

// get all posts
func (s *PostService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}

// get post by id
func (s *PostService) GetByID(id string) (*models.Post, error) {
	return s.repo.GetByID(id)
}

// delete post
func (s *PostService) Delete(id string) error {
	return s.repo.DeletePost(id)
}

// update a single post
func (s *PostService) Update(postID string, title string, content string) (*models.Post, error) {
	post, err := s.repo.GetByID(postID)
	if err != nil {
		return nil, err
	}

	post.Title = title
	post.Content = content
	post.Slug = utils.MakeSlugSimple(title + "-" + uuid.NewString()[:6])
	post.UpdatedAt = time.Now()

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&post).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}
