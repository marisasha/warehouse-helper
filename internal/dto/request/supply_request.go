package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Supplier struct {
	WarehouseID uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Contact     *string   `json:"contact,omitempty"`
	Phone       *string   `json:"phone,omitempty"`
	Address     *string   `json:"address,omitempty"`
}

type Supply struct {
	SupplierID uuid.UUID `json:"-"`
	CreatorID  uint64    `json:"-"`
	ProductID  uuid.UUID `json:"product_id"`
	Comment    *string   `json:"comment,omitempty"`

	Quantity    float64         `json:"quantity"`
	CostPerUnit decimal.Decimal `json:"cost_per_unit"`
	HarvestDate *time.Time      `json:"harvest_date,omitempty"`
	ExpiryDate  *time.Time      `json:"expiry_date,omitempty"`

	AcceptedAt *time.Time `json:"-"`
	CreatedAt  time.Time  `json:"-"`
}
