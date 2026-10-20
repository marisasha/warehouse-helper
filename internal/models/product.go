package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductCategoryType string
type ProductUnit string

const (
	TypeVegetables ProductCategoryType = "vegetables"
	TypeFruits     ProductCategoryType = "fruits"
	TypeBerries    ProductCategoryType = "berries"
)
const (
	UnitKg    ProductUnit = "kg"
	UnitPcs   ProductUnit = "pcs"
	UnitBox   ProductUnit = "box"
	UnitLiter ProductUnit = "liter"
)

type ProductCategory struct {
	ID          uuid.UUID           `json:"id"`
	WarehouseID uuid.UUID           `json:"warehouse_id"`
	Name        string              `json:"name"`
	Type        ProductCategoryType `json:"type"`

	Products  []Product `gorm:"foreignKey:CategoryID" json:"-"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID" json:"-"`
}

type Product struct {
	ID           uuid.UUID   `json:"id"`
	CategoryID   uuid.UUID   `json:"category_id"`
	SKU          string      `json:"sku"`
	Name         string      `json:"name"`
	Unit         ProductUnit `json:"unit"`
	PricePerUnit float64     `json:"price_per_unit"`
	Description  *string     `json:"description,omitempty"`
	MinStock     float64     `json:"min_stock"`
	IsActive     bool        `json:"is_active"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`

	// Связи
	Category ProductCategory `gorm:"foreignKey:CategoryID" json:"category"`
	Batches  []Batch         `gorm:"foreignKey:ProductID" json:"-"`
}

type Batch struct {
	ID          uuid.UUID       `json:"id"`
	ProductID   uuid.UUID       `json:"product_id"`
	Quantity    float64         `json:"quantity"`
	Remaining   float64         `json:"remaining"`
	CostPerUnit decimal.Decimal `json:"cost_per_unit"`
	HarvestDate *time.Time      `json:"harvest_date,omitempty"`
	ExpiryDate  *time.Time      `json:"expiry_date,omitempty"`
	ReceivedAt  time.Time       `json:"received_at"`

	// Связи
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}
