package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type SupplyStatus string

const (
	SupplyDraft     SupplyStatus = "draft"
	SupplyAccepted  SupplyStatus = "accepted"
	SupplyCancelled SupplyStatus = "cancelled"
)

type Supplier struct {
	ID          uuid.UUID `json:"id"`
	WarehouseID uuid.UUID `json:"warehouse_id"`
	Name        string    `json:"name"`
	Contact     *string   `json:"contact,omitempty"`
	Phone       *string   `json:"phone,omitempty"`
	Address     *string   `json:"address,omitempty"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`

	Warehouse Warehouse `gorm:"foreignKey:WarehouseID" json:"-"`
	Supplies  []Supply  `gorm:"foreignKey:SupplierID" json:"-"`
}

type Supply struct {
	ID         uuid.UUID    `json:"id"`
	SupplierID uuid.UUID    `json:"supplier_id"`
	ProductID  uuid.UUID    `json:"product_id"`
	CreatorID  uint         `json:"creator_id"`
	Status     SupplyStatus `json:"status"`
	Comment    *string      `json:"comment,omitempty"`

	Quantity    float64         `json:"quantity"`
	CostPerUnit decimal.Decimal `json:"cost_per_unit"`
	HarvestDate *time.Time      `json:"harvest_date,omitempty"`
	ExpiryDate  *time.Time      `json:"expiry_date,omitempty"`

	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`

	// Связи
	Product  Product  `gorm:"foreignKey:ProductID" json:"-"`
	Supplier Supplier `gorm:"foreignKey:SupplierID" json:"supplier"`
	Creator  User     `gorm:"foreignKey:CreatorID" json:"-"`
}
