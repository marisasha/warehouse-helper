package repository

import (
	"errors"

	dto "github.com/marisasha/warehouse-helper/internal/dto/request"
	"gorm.io/gorm"
)

type AuthDB struct {
	db *gorm.DB
}

func NewAuthDB(db *gorm.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(user *dto.User) error {
	return r.db.Create(user).Error
}

func (r *AuthDB) GetUser(email, passwordHash string) (dto.User, error) {
	var user dto.User

	err := r.db.
		Where("email = ? AND password_hash = ?", email, passwordHash).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.User{}, err
	}

	return user, err
}
