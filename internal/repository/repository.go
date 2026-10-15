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
	GetWarehouseEpmloyees(warehouseID uuid.UUID) ([]*models.WarehouseEmployee, error)
	GetWarehouseEmployee(warehouseID, employeeID uuid.UUID) (*models.WarehouseEmployee, error)
	DeleteWarehouseEmployee(warehouseID, employeeID uuid.UUID) error
}

type Product interface {
	CreateProductCategory(input *req.ProductCategory) (uuid.UUID, error)
	GetProductCategories(warehouseID uuid.UUID) ([]*models.ProductCategory, error)
	GetProductCategory(warehouseID, categoryID uuid.UUID) (*models.ProductCategory, error)
	DeleteProductCategory(warehouseID, categoryID uuid.UUID) error

	CreateProduct(input *req.Product) (uuid.UUID, error)
	GetProductsByCategory(warehouseID, categoryID uuid.UUID) ([]*models.Product, error)
	GetProduct(warehouseID, categoryID, productID uuid.UUID) (*models.Product, error)
	DeleteProduct(warehouseID, categoryID, productID uuid.UUID) error
}

type Repository struct {
	Authorization
	Warehouse
	Product
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Warehouse:     NewWarehouseDB(db),
		Product:       NewProductDB(db),
	}
}
