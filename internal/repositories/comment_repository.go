package repositories

import (
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Create(c *models.Comment) error {
	return r.db.Create(c).Error
}

func (r *CommentRepository) ListByPost(postId string) ([]models.Comment, error) {

	var cs []models.Comment

	if err := r.db.Where("post_id=?", postId).Find(&cs).Error; err != nil {
		return nil, err

	}
	return cs, nil
}

func (r *CommentRepository) GetByID(id string) (*models.Comment, error) {
	var c models.Comment
	if err := r.db.Where("id=?", id).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CommentRepository) Update(c *models.Comment) error {
	return r.db.Save(c).Error
}

func (r *CommentRepository) Delete(id string) error {
	return r.db.Delete(&models.Comment{}, "id=?", id).Error
}
