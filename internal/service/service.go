package service

import (
	dto "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/repository"
)

type Authorization interface {
	CreateUser(user *dto.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token *string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository, cfg Argon2id) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, cfg),
	}
}
