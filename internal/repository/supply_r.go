package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SupplyDB struct {
	db *gorm.DB
}

func NewSupplyDB(db *gorm.DB) *SupplyDB {
	return &SupplyDB{db: db}
}

func (r *SupplyDB) CreateSupplier(input *req.Supplier) (uuid.UUID, error) {
	var created *models.Supplier

	err := r.db.
		Table(supplierTable).
		Clauses(clause.Returning{}).
		Create(input).
		Scan(&created).Error

	if err != nil {
		return uuid.Nil, err
	}

	return created.ID, nil
}

func (r *SupplyDB) GetSuppliers(warehouseID uuid.UUID) ([]*models.Supplier, error) {
	var suppliers []*models.Supplier

	result := r.db.
		Table(supplierTable).
		Where("warehouse_id = ?", warehouseID).
		Find(&suppliers)
	if result.Error != nil {
		return nil, fmt.Errorf("error in repository: %s", result.Error)
	}
	return suppliers, nil

}
func (r *SupplyDB) GetSupplier(warehouseID, supplierID uuid.UUID) (*models.Supplier, error) {
	var supplier *models.Supplier

	result := r.db.
		Table(supplierTable).
		Where("warehouse_id = ? AND id = ?", warehouseID, supplierID).
		First(&supplier)
	if result.Error != nil {
		return nil, fmt.Errorf("error in repository: %s", result.Error)
	}
	return supplier, nil
}
func (r *SupplyDB) GetSupplierExist(warehouseID, supplierID uuid.UUID) (bool, error) {
	var count int64

	result := r.db.
		Table(supplierTable).
		Where("warehouse_id = ? AND id = ?", warehouseID, supplierID).
		Count(&count)

	if result.Error != nil {
		return false, fmt.Errorf("error in repository: %s", result.Error)
	}
	return count > 0, nil
}

func (r *SupplyDB) DeleteSupplier(warehouseID, supplierID uuid.UUID) error {
	result := r.db.
		Table(supplierTable).
		Where("warehouse_id = ? AND id = ?", warehouseID, supplierID).
		Delete(nil)

	if result.Error != nil {
		return fmt.Errorf("error in repository: %s", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Продукт с id %s не найден", supplierID)
	}

	return nil
}

func (r *SupplyDB) CreateSupply(input *req.Supply) (uuid.UUID, error) {
	var created *models.Supply

	result := r.db.
		Table(supplyTable).
		Clauses(clause.Returning{}).
		Create(input).
		Scan(&created)

	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("error in repository: %s", result.Error)
	}

	return created.ID, nil
}

func (r *SupplyDB) GetEmployeeExist(warehouseID uuid.UUID, userID uint64) (bool, error) {
	var count int64

	result := r.db.
		Table(warehouseEmployeesTable).
		Where("warehouse_id = ? AND user_id = ?", warehouseID, userID).
		Count(&count)

	if result.Error != nil {
		return false, fmt.Errorf("error in repository: %s", result.Error)
	}
	return count > 0, nil
}

func (r *SupplyDB) UpdateStatus(ctx context.Context, input *req.SupplyStatus, supplyID uuid.UUID) (*models.Supply, error) {

	var supply *models.Supply

	if err := r.db.WithContext(ctx).
		Table(supplyTable).
		Where("id = ?", supplyID).
		First(supply).Error; err != nil {
		return nil, fmt.Errorf("error in repository: %s", err)
	}

	if supply.Status != models.SupplyDraft {
		return nil, fmt.Errorf("error in repository: можно обновить поставку только со статусом draft, текущий статус: %s", supply.Status)
	}

	result := r.db.
		WithContext(ctx).
		Table(supplyTable).
		Where("id = ?", supplyID).
		Updates(map[string]any{
			"status":          input.Status,
			"last_updator_id": input.LastUpdatorID,
		})

	if result.Error != nil {
		return nil, fmt.Errorf("error in repository: %s", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return supply, nil
}

func (r *SupplyDB) CreateBatch(ctx context.Context, supply *models.Supply) (uuid.UUID, error) {
	created := models.Batch{
		ID:          uuid.New(),
		ProductID:   supply.ProductID,
		Quantity:    supply.Quantity,
		Remaining:   supply.Quantity,
		CostPerUnit: supply.CostPerUnit,
		HarvestDate: supply.HarvestDate,
		ExpiryDate:  supply.ExpiryDate,
		ReceivedAt:  time.Now(),
	}

	result := r.db.
		WithContext(ctx).
		Table(batchesTable).
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Create(&created)

	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("error in repository: %s", result.Error)
	}
	return created.ID, nil
}
