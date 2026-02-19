package postgres

import (
	"github.com/AfshinNajafi74/go-gymApp/internal/domain/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(u *user.User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) GetByEmail(email string) (*user.User, error) {
	var u user.User
	err := r.db.Where("email = ?", email).First(&u).Error
	return &u, err
}
