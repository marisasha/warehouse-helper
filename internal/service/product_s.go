package service

import (
	"fmt"

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
	return s.repos.GetProductCategoryData(warehouseID, categoryID)
}

func (s *ProductService) DeleteProductCategory(warehouseID, categoryID uuid.UUID) error {
	return s.repos.DeleteProductCategory(warehouseID, categoryID)
}

func (s *ProductService) CreateProduct(input *req.Product, warehouseID uuid.UUID) (uuid.UUID, error) {
	isCategoryExist, err := s.repos.GetProductCategoryExist(warehouseID, input.CategoryID)
	if err != nil {
		return uuid.Nil, err
	}
	if !isCategoryExist {
		return uuid.Nil, fmt.Errorf("error in service: Не существует категории %s с warehouseID %s", input.CategoryID, warehouseID)
	}
	return s.repos.CreateProduct(input)
}

func (s *ProductService) GetProductsByCategory(warehouseID, categoryID uuid.UUID) ([]*models.Product, error) {
	isCategoryExist, err := s.repos.GetProductCategoryExist(warehouseID, categoryID)
	if err != nil {
		return nil, err
	}
	if !isCategoryExist {
		return nil, fmt.Errorf("error in service: Не существует категории %s с warehouseID %s", categoryID, warehouseID)
	}
	return s.repos.GetProductsByCategory(categoryID)
}

func (s *ProductService) GetProduct(warehouseID, categoryID, productID uuid.UUID) (*models.Product, error) {
	isCategoryExist, err := s.repos.GetProductCategoryExist(warehouseID, categoryID)
	if err != nil {
		return nil, err
	}
	if !isCategoryExist {
		return nil, fmt.Errorf("error in service: Не существует категории %s с warehouseID %s", categoryID, warehouseID)
	}
	return s.repos.GetProduct(categoryID, productID)
}

func (s *ProductService) DeleteProduct(warehouseID, categoryID, productID uuid.UUID) error {
	isCategoryExist, err := s.repos.GetProductCategoryExist(warehouseID, categoryID)
	if err != nil {
		return err
	}
	if !isCategoryExist {
		return fmt.Errorf("error in service: Не существует категории %s с warehouseID %s", categoryID, warehouseID)
	}
	return s.repos.DeleteProduct(categoryID, productID)
}
