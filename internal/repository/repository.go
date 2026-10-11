package repository

import (
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *req.User) (uint64, error)
	GetUser(email string) (uint64, string, error)
}

type Warehouse interface {
	GetWarehouseDetail(warehouseID *uuid.UUID) (*models.Warehouse, error)
	CreateWarhouseEmployee(input *req.WarehouseEmployee) (uuid.UUID, error)
}

type Repository struct {
	Authorization
	Warehouse
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Warehouse:     NewWarehouseDB(db),
	}
}
