package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductUnit string
type BatchStatus string

const (
	UnitKg    ProductUnit = "kg"
	UnitPcs   ProductUnit = "pcs"
	UnitBox   ProductUnit = "box"
	UnitLiter ProductUnit = "liter"
)
const (
	BatchActive     BatchStatus = "active"
	BatchExpired    BatchStatus = "expired"
	BatchWrittenOff BatchStatus = "written_off"
)

type ProductCategory struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Type string    `json:"type"`

	Products []Product `gorm:"foreignKey:CategoryID" json:"-"`
}

type Product struct {
	ID           uuid.UUID   `json:"id"`
	CategoryID   uuid.UUID   `json:"category_id"`
	WarehouseID  uuid.UUID   `json:"warehouse_id"`
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
	Warehouse Warehouse       `gorm:"foreignKey:WarehouseID" json:"-"`
	Category  ProductCategory `gorm:"foreignKey:CategoryID" json:"category"`
	Batches   []Batch         `gorm:"foreignKey:ProductID" json:"-"`
}

type Batch struct {
	ID           uuid.UUID       `json:"id"`
	WarehouseID  uuid.UUID       `json:"warehouse_id"`
	ProductID    uuid.UUID       `json:"product_id"`
	SupplyItemID uuid.UUID       `json:"supply_item_id"`
	Quantity     float64         `json:"quantity"`
	Remaining    float64         `json:"remaining"`
	CostPerUnit  decimal.Decimal `json:"cost_per_unit"`
	HarvestDate  *time.Time      `json:"harvest_date,omitempty"`
	ExpiryDate   *time.Time      `json:"expiry_date,omitempty"`
	ReceivedAt   time.Time       `json:"received_at"`
	Status       BatchStatus     `json:"status"`

	// Связи
	Warehouse  Warehouse   `gorm:"foreignKey:WarehouseID" json:"-"`
	Product    Product     `gorm:"foreignKey:ProductID" json:"product"`
	SupplyItem *SupplyItem `gorm:"foreignKey:SupplyItemID" json:"-"`
}
