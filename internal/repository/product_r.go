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

	result := r.db.
		Table(productCategoriesTable).
		Clauses(clause.Returning{}).
		Create(&input).
		Scan(&created)

	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("error in repository: %s", result.Error)
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
		return nil, fmt.Errorf("error in repository: %s", result.Error)
	}

	return categories, nil
}

func (r *ProductDB) GetProductCategoryData(warehouseID, categoryID uuid.UUID) (*models.ProductCategory, error) {
	var category *models.ProductCategory

	result := r.db.
		Table(productCategoriesTable).
		Where("warehouse_id = ? AND id = ?", warehouseID, categoryID).
		First(&category)

	if result.Error != nil {
		return nil, fmt.Errorf("error in repository: %s", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Категория с id %s не найдена", categoryID)
	}

	return category, nil
}
func (r *ProductDB) GetProductCategoryExist(warehouseID, categoryID uuid.UUID) (bool, error) {
	var count int64

	result := r.db.
		Table(productCategoriesTable).
		Where("warehouse_id = ? AND id = ?", warehouseID, categoryID).
		Count(&count)

	if result.Error != nil {
		return false, fmt.Errorf("error in repository: %s", result.Error)
	}
	return count > 0, nil
}

func (r *ProductDB) DeleteProductCategory(warehouseID, categoryID uuid.UUID) error {
	result := r.db.
		Table(productCategoriesTable).
		Where("warehouse_id = ? AND id = ?", warehouseID, categoryID).
		Delete(nil)

	if result.Error != nil {
		return fmt.Errorf("error in repository: %s", result.Error)
	}

	return nil
}

func (r *ProductDB) CreateProduct(input *req.Product) (uuid.UUID, error) {
	var created *models.Product

	result := r.db.
		Table(productTable).
		Clauses(clause.Returning{}).
		Create(&input).
		Scan(&created)

	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("error in repository: %s", result.Error)
	}

	return created.ID, nil
}
func (r *ProductDB) GetProductsByCategory(categoryID uuid.UUID) ([]*models.Product, error) {
	var products []*models.Product
	result := r.db.
		Table(productTable).
		Where("category_id = ?", categoryID).
		Find(&products)

	if result.Error != nil {
		return nil, fmt.Errorf("error in repository: %s", result.Error)
	}

	return products, nil
}
func (r *ProductDB) GetProduct(categoryID, productID uuid.UUID) (*models.Product, error) {
	var product *models.Product

	result := r.db.
		Table(productTable).
		Where("category_id = ? AND id = ?", categoryID, productID).
		First(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}
func (r *ProductDB) DeleteProduct(categoryID, productID uuid.UUID) error {
	result := r.db.
		Table(productTable).
		Where("category_id = ? AND id = ?", categoryID, productID).
		Delete(nil)

	if result.Error != nil {
		return fmt.Errorf("error in repository: %s", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Продукт с id %s не найден", productID)
	}

	return nil
}
