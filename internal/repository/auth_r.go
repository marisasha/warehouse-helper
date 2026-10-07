package repository

import (
	"errors"

	dto "github.com/marisasha/warehouse-helper/internal/dto/request"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user *dto.User) error {
	return r.db.Create(user).Error
}

func (r *AuthPostgres) GetUser(email, passwordHash string) (dto.User, error) {
	var user dto.User

	err := r.db.
		Where("email = ? AND password_hash = ?", email, passwordHash).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.User{}, err
	}

	return user, err
}
