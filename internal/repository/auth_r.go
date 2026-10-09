package repository

import (
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

func (r *AuthDB) GetUser(email string) (int, string, error) {
	var user dto.User

	err := r.db.
		Table(userTable).
		Where("email = ?", email).
		First(&user).Error

	return user.ID, user.Password, err
}
