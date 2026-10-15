package request

import (
	"github.com/google/uuid"
	"github.com/marisasha/warehouse-helper/internal/models"
)

type ProductCategory struct {
	WarehouseID uuid.UUID                  `json:"-"`
	Name        string                     `json:"name"`
	Type        models.ProductCategoryType `json:"type"`
}

type Product struct {
	CategoryID   uuid.UUID          `json:"-"`
	SKU          string             `json:"sku" default:""`
	Name         string             `json:"name"`
	Unit         models.ProductUnit `json:"unit"`
	PricePerUnit float64            `json:"price_per_unit"`
	Description  *string            `json:"description,omitempty"`
	MinStock     float64            `json:"min_stock"`
}

// f84b3e53-ef95-45dc-8217-eeb9586a3cf8
// 74469cee-f96b-44ad-8fa4-6f250d34aa54

// {
//   "description": "вкусные помидоры ",
//   "min_stock": 1,
//   "name": "Помидоры Топар",
//   "price_per_unit": 100,
//   "sku": "123456",
//   "unit": "kg"
// }
