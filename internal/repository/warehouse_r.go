package repository

import (
	"fmt"

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

func (r *WarehouseDB) GetWarehouseEpmloyees(warehouseID uuid.UUID) ([]*models.WarehouseEmployee, error) {
	var warehouseEmployees []*models.WarehouseEmployee

	result := r.db.
		Table(warehouseEmployeesTable).
		Where("warehouse_id = ?", warehouseID).
		Find(&warehouseEmployees)

	if result.Error != nil {
		return nil, result.Error
	}

	return warehouseEmployees, nil
}

func (r *WarehouseDB) GetWarehouseEmployee(warehouseID, employeeID uuid.UUID) (*models.WarehouseEmployee, error) {
	var warehouseEmployees *models.WarehouseEmployee

	result := r.db.
		Table(warehouseEmployeesTable).
		Where("warehouse_id = ? and id = ?", warehouseID, employeeID).
		First(&warehouseEmployees)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Сотрудник с id %s не найден", employeeID)
	}

	return warehouseEmployees, nil
}
func (r *WarehouseDB) DeleteWarehouseEmployee(warehouseID, employeeID uuid.UUID) error {

	result := r.db.
		Table(warehouseEmployeesTable).
		Where("warehouse_id = ? AND id = ?", warehouseID, employeeID).
		Delete(nil)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Сотрудник с id %s не найден", employeeID)
	}

	return nil
}
