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
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Contact   *string   `json:"contact,omitempty"`
	Phone     *string   `json:"phone,omitempty"`
	Address   *string   `json:"address,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`

	Supplies []Supply `gorm:"foreignKey:SupplierID" json:"-"`
}

type Supply struct {
	ID          uuid.UUID    `json:"id"`
	WarehouseID uuid.UUID    `json:"warehouse_id"`
	SupplierID  uuid.UUID    `json:"supplier_id"`
	CreatorID   uint         `json:"creator_id"`
	Status      SupplyStatus `json:"status"`
	AcceptedAt  *time.Time   `json:"accepted_at,omitempty"`
	Comment     *string      `json:"comment,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`

	// Связи
	Warehouse Warehouse    `gorm:"foreignKey:WarehouseID" json:"-"`
	Supplier  Supplier     `gorm:"foreignKey:SupplierID" json:"supplier"`
	Creator   User         `gorm:"foreignKey:CreatorID" json:"-"`
	Items     []SupplyItem `gorm:"foreignKey:SupplyID" json:"items"`
}

type SupplyItem struct {
	ID          uuid.UUID       `json:"id"`
	SupplyID    uuid.UUID       `json:"supply_id"`
	ProductID   uuid.UUID       `json:"product_id"`
	BatchID     *uuid.UUID      `json:"batch_id,omitempty"`
	Quantity    float64         `json:"quantity"`
	CostPerUnit decimal.Decimal `json:"cost_per_unit"`
	HarvestDate *time.Time      `json:"harvest_date,omitempty"`
	ExpiryDate  *time.Time      `json:"expiry_date,omitempty"`

	Supply  Supply  `gorm:"foreignKey:SupplyID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
	Batch   *Batch  `gorm:"foreignKey:BatchID" json:"-"`
}
