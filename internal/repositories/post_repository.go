package repositories

import (
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

// create a single post

func (r *PostRepository) Create(p *models.Post) error {
	return r.db.Create(p).Error
}

// get all posts
func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Order("created_at desc").Find(&posts).Error
	return posts, err
}

// get post by id

func (r *PostRepository) GetByID(id string) (*models.Post, error) {
	var post models.Post
	err := r.db.Preload("Comments").Preload("Reactions").Preload("Tags").Where("id=?", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// delete a post

func (r *PostRepository) DeletePost(id string) error {

	return r.db.Transaction(func(tx *gorm.DB) error {
		var post models.Post
		if err := tx.Preload("Tags").First(&post, "id=?", id).Error; err != nil {
			return nil
		}
		if err := tx.Model(&post).Association("Tags").Clear(); err != nil {
			return nil
		}
		if err := tx.Delete(&post).Error; err != nil {
			return nil
		}
		return nil
	})
}

// update a post

func (r *PostRepository) UpdatePost(post *models.Post) error {
	return r.db.Save(post).Error
}
