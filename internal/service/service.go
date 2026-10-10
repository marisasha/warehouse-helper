package service

import (
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/marisasha/warehouse-helper/internal/repository"
)

type Authorization interface {
	CreateUser(user *req.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token *string) (int, error)
}
type Warehouse interface {
	GetWarehouseDetail(warehouseId *uuid.UUID) (*models.Warehouse, error)
}

type Service struct {
	Authorization
	Warehouse
}

func NewService(repos *repository.Repository, cfg Argon2id) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, cfg),
		Warehouse:     NewWarehouseService(repos.Warehouse, cfg),
	}
}
