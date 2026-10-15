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
