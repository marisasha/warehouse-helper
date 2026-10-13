package service

import (
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/marisasha/warehouse-helper/internal/repository"
)

type WarehouseService struct {
	repos repository.Warehouse
}

func NewWarehouseService(repos repository.Warehouse) *WarehouseService {
	return &WarehouseService{
		repos: repos,
	}
}

func (s *WarehouseService) GetWarehouseDetail(warehouseID *uuid.UUID) (*models.Warehouse, error) {
	return s.repos.GetWarehouseDetail(warehouseID)
}

func (s *WarehouseService) CreateWarhouseEmployee(input *req.WarehouseEmployee) (uuid.UUID, error) {
	return s.repos.CreateWarhouseEmployee(input)
}

func (s *WarehouseService) GetWarehouseEpmloyees(warehouseID uuid.UUID) ([]*models.WarehouseEmployee, error) {
	return s.repos.GetWarehouseEpmloyees(warehouseID)
}

func (s *WarehouseService) GetWarehouseEmployee(warehouseID, employeeID uuid.UUID) (*models.WarehouseEmployee, error) {
	return s.repos.GetWarehouseEmployee(warehouseID, employeeID)

}

func (s *WarehouseService) DeleteWarehouseEmployee(warehouseID, employeeID uuid.UUID) error {
	return s.repos.DeleteWarehouseEmployee(warehouseID, employeeID)
}
