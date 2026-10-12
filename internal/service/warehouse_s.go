package service

import (
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/marisasha/warehouse-helper/internal/repository"
)

type WarehouseService struct {
	repos repository.Warehouse
	cfg   Argon2id
}

func NewWarehouseService(repos repository.Warehouse, cfg Argon2id) *WarehouseService {
	return &WarehouseService{
		repos: repos,
		cfg:   cfg,
	}
}

func (s *WarehouseService) GetWarehouseDetail(warehouseID *uuid.UUID) (*models.Warehouse, error) {
	warehouse, err := s.repos.GetWarehouseDetail(warehouseID)
	if err != nil {
		return nil, err
	}

	return warehouse, nil

}

func (s *WarehouseService) CreateWarhouseEmployee(input *req.WarehouseEmployee) (uuid.UUID, error) {
	warehouseEmployeeID, err := s.repos.CreateWarhouseEmployee(input)
	if err != nil {
		return uuid.Nil, err
	}

	return warehouseEmployeeID, nil
}

func (s *WarehouseService) GetWarehouseEpmloyees(warehouseID uuid.UUID) ([]*models.WarehouseEmployee, error) {
	warehouseEmployees, err := s.repos.GetWarehouseEpmloyees(warehouseID)
	if err != nil {
		return nil, err
	}
	return warehouseEmployees, nil
}

func (s *WarehouseService) GetWarehouseEmployee(warehouseID, employeeID uuid.UUID) (*models.WarehouseEmployee, error) {
	warehouseEmployee, err := s.repos.GetWarehouseEmployee(warehouseID, employeeID)
	if err != nil {
		return nil, err
	}
	return warehouseEmployee, nil
}

func (s *WarehouseService) DeleteWarehouseEmployee(warehouseID, employeeID uuid.UUID) error {
	err := s.repos.DeleteWarehouseEmployee(warehouseID, employeeID)
	if err != nil {
		return err
	}
	return nil
}
