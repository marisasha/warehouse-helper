package repository

import (
	dto "github.com/marisasha/warehouse-helper/internal/dto/request"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *dto.User) error
	GetUser(email, passwordHash string) (dto.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
	}
}
