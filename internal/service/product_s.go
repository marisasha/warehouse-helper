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
