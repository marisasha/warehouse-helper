package service

import (
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/marisasha/warehouse-helper/internal/repository"
)

type SupplyService struct {
	repos repository.Supply
}

func NewSupplyService(repos repository.Supply) *SupplyService {
	return &SupplyService{
		repos: repos,
	}
}

func (s *SupplyService) CreateSupplier(input *req.Supplier) (uuid.UUID, error) {
	return s.repos.CreateSupplier(input)
}

func (s *SupplyService) GetSuppliers(warehouseID uuid.UUID) ([]*models.Supplier, error) {
	return s.repos.GetSuppliers(warehouseID)
}

func (s *SupplyService) GetSupplier(warehouseID, supplierID uuid.UUID) (*models.Supplier, error) {
	return s.repos.GetSupplier(warehouseID, supplierID)
}

func (s *SupplyService) DeleteSupplier(warehouseID, supplierID uuid.UUID) error {
	return s.repos.DeleteSupplier(warehouseID, supplierID)
}
