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
		return uuid.Nil, result.Error
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
		First(category)

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
		return uuid.Nil, result.Error
	}

	return created.ID, nil
}
func (r *ProductDB) GetProductsByCategory(warehouseID, categoryID uuid.UUID) ([]*models.Product, error) {
	var products []*models.Product

	result := r.db.
		Table(productTable).
		Where("warehouse_id = ? and category_id = ?", warehouseID, categoryID).
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
func (r *ProductDB) GetProduct(warehouseID, categoryID, productID uuid.UUID) (*models.Product, error) {
	var product *models.Product

	result := r.db.
		Table(usersTable).
		Where("warehouse_id = ? and category_id = ? and id = ?", warehouseID, categoryID, productID).
		First(product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}
func (r *ProductDB) DeleteProduct(warehouseID, categoryID, productID uuid.UUID) error {
	result := r.db.
		Table(usersTable).
		Where("warehouse_id = ? and category_id = ? and id = ?", warehouseID, categoryID, productID).
		Delete(nil)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Продукт с id %s не найден", productID)
	}

	return nil
}
