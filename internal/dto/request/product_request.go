package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/marisasha/warehouse-helper/internal/models"
	"github.com/shopspring/decimal"
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

type Batch struct {
	ProductID    uuid.UUID          `json:"product_id"`
	SupplyItemID uuid.UUID          `json:"supply_item_id"`
	Quantity     float64            `json:"quantity"`
	Remaining    float64            `json:"remaining"`
	CostPerUnit  decimal.Decimal    `json:"cost_per_unit"`
	HarvestDate  *time.Time         `json:"harvest_date,omitempty"`
	ExpiryDate   *time.Time         `json:"expiry_date,omitempty"`
	Status       models.BatchStatus `json:"status"`
}
