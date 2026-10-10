package service

import (
	"github.com/google/uuid"
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
