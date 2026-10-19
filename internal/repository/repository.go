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
	GetProductCategoryData(warehouseID, categoryID uuid.UUID) (*models.ProductCategory, error)
	GetProductCategoryExist(warehouseID, categoryID uuid.UUID) (bool, error)
	DeleteProductCategory(warehouseID, categoryID uuid.UUID) error

	CreateProduct(input *req.Product) (uuid.UUID, error)
	GetProductsByCategory(categoryID uuid.UUID) ([]*models.Product, error)
	GetProduct(categoryID, productID uuid.UUID) (*models.Product, error)
	DeleteProduct(categoryID, productID uuid.UUID) error
}
type Supply interface {
	CreateSupplier(input *req.Supplier) (uuid.UUID, error)
	GetSuppliers(warehouseID uuid.UUID) ([]*models.Supplier, error)
	GetSupplier(warehouseID, supplierID uuid.UUID) (*models.Supplier, error)
	GetSupplierExist(warehouseID, supplierID uuid.UUID) (bool, error)
	DeleteSupplier(warehouseID, supplierID uuid.UUID) error

	CreateSupply(input *req.Supply) (uuid.UUID, error)
	GetEmployeeExist(warehouseID uuid.UUID, userID uint64) (bool, error)
	UpdateStatus(input *req.SupplyStatus, supplyID uuid.UUID) (*models.Supply, error)
}
type Repository struct {
	Authorization
	Warehouse
	Product
	Supply
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Warehouse:     NewWarehouseDB(db),
		Product:       NewProductDB(db),
		Supply:        NewSupplyDB(db),
	}
}
