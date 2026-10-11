package repository

import (
	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		First(&warehouse).Error
	if err != nil {
		return nil, err
	}

	return warehouse, nil
}

func (r *WarehouseDB) CreateWarhouseEmployee(input *req.WarehouseEmployee) (uuid.UUID, error) {
	var created *models.WarehouseEmployee

	err := r.db.
		Table(warehouseEmployeesTable).
		Clauses(clause.Returning{}).
		Create(input).
		Scan(&created).Error

	if err != nil {
		return uuid.Nil, err
	}

	return created.ID, nil

}
