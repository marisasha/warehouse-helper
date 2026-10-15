package service

import (
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/marisasha/warehouse-helper/internal/repository"
)

type ProductService struct {
	repos repository.Product
}

func NewProductService(repos repository.Product) *ProductService {
	return &ProductService{
		repos: repos,
	}
}

func (s *ProductService) CreateProductCategory(input *req.ProductCategory) (uuid.UUID, error) {
	return s.repos.CreateProductCategory(input)
}

func (s *ProductService) GetProductCategories(warehouseID uuid.UUID) ([]*models.ProductCategory, error) {
	return s.repos.GetProductCategories(warehouseID)
}

func (s *ProductService) GetProductCategory(warehouseID, categoryID uuid.UUID) (*models.ProductCategory, error) {
	return s.repos.GetProductCategory(warehouseID, categoryID)
}

func (s *ProductService) DeleteProductCategory(warehouseID, categoryID uuid.UUID) error {
	return s.repos.DeleteProductCategory(warehouseID, categoryID)
}

func (s *ProductService) CreateProduct(input *req.Product) (uuid.UUID, error) {
	return s.repos.CreateProduct(input)
}

func (s *ProductService) GetProductsByCategory(warehouseID, categoryID uuid.UUID) ([]*models.Product, error) {
	return s.repos.GetProductsByCategory(warehouseID, categoryID)
}

func (s *ProductService) GetProduct(warehouseID, categoryID, productID uuid.UUID) (*models.Product, error) {
	return s.repos.GetProduct(warehouseID, categoryID, productID)
}

func (s *ProductService) DeleteProduct(warehouseID, categoryID, productID uuid.UUID) error {
	return s.repos.DeleteProduct(warehouseID, categoryID, productID)
}
