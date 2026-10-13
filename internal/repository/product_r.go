package repository

import (
	"fmt"

	"github.com/google/uuid"
	req "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductDB struct {
	db *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{db: db}
}
func (r *ProductDB) CreateProductCategory(input *req.ProductCategory) (uuid.UUID, error) {
	var created *models.ProductCategory

	err := r.db.
		Table(productCategoriesTable).
		Clauses(clause.Returning{}).
		Create(&input).
		Scan(&created).Error

	if err != nil {
		return uuid.Nil, err
	}

	return created.ID, nil
}

func (r *ProductDB) GetProductCategories(warehouseID uuid.UUID) ([]*models.ProductCategory, error) {
	var categories []*models.ProductCategory

	result := r.db.
		Table(productCategoriesTable).
		Where("warehouse_id = ?", warehouseID).
		Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (r *ProductDB) GetProductCategory(warehouseID, categoryID uuid.UUID) (*models.ProductCategory, error) {
	var category *models.ProductCategory

	result := r.db.
		Table(productCategoriesTable).
		Where("warehouse_id = ? and id = ?", warehouseID, categoryID).
		First(&category)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Категория с id %s не найдена", categoryID)
	}

	return category, nil
}

func (r *ProductDB) DeleteProductCategory(warehouseID, categoryID uuid.UUID) error {
	result := r.db.
		Table(productCategoriesTable).
		Where("warehouse_id = ? and id = ?", warehouseID, categoryID).
		Delete(nil)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Категория с id %s не найдена", categoryID)
	}

	return nil
}
