package repositories

import (
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u *models.User) error {
	return r.db.Create(u).Error
}

// get user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var u models.User
	if err := r.db.Where("email=?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
