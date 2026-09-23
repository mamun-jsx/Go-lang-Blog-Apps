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

func newPostService(r *repositories.PostRepository, db *gorm.DB) *PostService {
	return &PostService{
		repo: r,
		db:   db,
	}
}


// create a post 
func (s *PostService) Create(authorID uuid.UUID, title, content string, tags []string) (*models.Post, error) {

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
	err := s.db.Transaction(func(tx *gorm.DB) err {
		if err := tx.Create(post).Error; err != nil {
			return err
		}
		for _, tagName := range tags {
			var tag models.Tag
			if err := tx.Where("name=?", tagName).First(&tag).Error; err != nil {
				tag = models.Tag{
					ID:        uuid.New(),
					Name:      tagName,
					Slug:      utils.MakeSlugSimple(tagName),
					CreatedAt: now,
					UpdatedAt: now,
				}
				if err := tx.Create(&tag).Error; err != nil {
					return err
				}
				if err := tx.Model(post).Association("Tags").Append(&tag); err != nil {
					return nil
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return post, nil
}
