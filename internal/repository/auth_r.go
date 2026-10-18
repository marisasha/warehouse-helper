package repository

import (
	"errors"
	"fmt"

	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	res "github.com/marisasha/warehouse-helper/internal/dto/response"
	"github.com/marisasha/warehouse-helper/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthDB struct {
	db *gorm.DB
}

func NewAuthDB(db *gorm.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(user *req.User) (uint64, error) {
	var userModel models.User

	err := r.db.
		Table(usersTable).
		Clauses(clause.Returning{}).
		Create(user).
		Scan(&userModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return 0, fmt.Errorf("Пользователь с email %s уже существует", user.Email)
		}
		return 0, err
	}
	return userModel.ID, nil

}

func (r *AuthDB) GetUser(email string) (uint64, string, error) {
	var user res.UserResponse

	err := r.db.
		Table(usersTable).
		Where("email = ?", email).
		First(&user).Error

	return user.ID, user.PasswordHash, err
}
