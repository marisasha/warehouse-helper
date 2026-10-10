package repository

import (
	"errors"
	"fmt"

	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	res "github.com/marisasha/warehouse-helper/internal/dto/response"
	"gorm.io/gorm"
)

type AuthDB struct {
	db *gorm.DB
}

func NewAuthDB(db *gorm.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(user *req.User) error {
	err := r.db.
		Table(usersTable).
		Create(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("Пользователь с email %s уже существует", user.Email)
		}
		return err
	}
	return nil

}

func (r *AuthDB) GetUser(email string) (uint64, string, error) {
	var user res.UserResponse

	err := r.db.
		Table(usersTable).
		Where("email = ?", email).
		First(&user).Error

	return user.ID, user.Password, err
}
