package repository

import (
	"github.com/google/uuid"
	"github.com/marisasha/warehouse-helper/internal/models"
	"gorm.io/gorm"
)

type WarehouseDB struct {
	db *gorm.DB
}

func NewWarehouseDB(db *gorm.DB) *WarehouseDB {
	return &WarehouseDB{db: db}
}

func (r *WarehouseDB) GetWarehouseDetail(warehouseID *uuid.UUID) (*models.Warehouse, error) {
	var warehouse *models.Warehouse

	err := r.db.
		Table(warehousesTable).
		Where("id = ?", *warehouseID).
		First(warehouse).Error
	if err != nil {
		return nil, err
	}

	return warehouse, nil
}
