package request

import "github.com/google/uuid"

type ProductCategory struct {
	WarehouseID uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
}
