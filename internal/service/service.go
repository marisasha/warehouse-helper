package service

import (
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/marisasha/warehouse-helper/internal/repository"
)

type Authorization interface {
	CreateUser(user *req.User) (uint64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token *string) (int, error)
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

	CreateProduct(input *req.Product, warehouseID uuid.UUID) (uuid.UUID, error)
	GetProductsByCategory(warehouseID, categoryID uuid.UUID) ([]*models.Product, error)
	GetProduct(warehouseID, categoryID, productID uuid.UUID) (*models.Product, error)
	DeleteProduct(warehouseID, categoryID, productID uuid.UUID) error
}

type Supply interface {
	CreateSupplier(input *req.Supplier) (uuid.UUID, error)
	GetSuppliers(warehouseID uuid.UUID) ([]*models.Supplier, error)
	GetSupplier(warehouseID, supplierID uuid.UUID) (*models.Supplier, error)
	DeleteSupplier(warehouseID, supplierID uuid.UUID) error
}
type Service struct {
	Authorization
	Warehouse
	Product
	Supply
}

func NewService(repos *repository.Repository, cfg Argon2id) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, cfg),
		Warehouse:     NewWarehouseService(repos.Warehouse),
		Product:       NewProductService(repos),
		Supply:        NewSupplyService(repos),
	}
}
