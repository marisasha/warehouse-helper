package models

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeRole string

const (
	RoleOwner  EmployeeRole = "owner"
	RoleAdmin  EmployeeRole = "admin"
	RoleWorker EmployeeRole = "worker"
)

type Warehouse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   *string   `json:"address,omitempty"`
	Country   *string   `json:"country,omitempty"`
	Region    *string   `json:"region,omitempty"`
	District  *string   `json:"district,omitempty"`
	City      *string   `json:"city,omitempty"`
	Timezone  string    `json:"timezone"`
	Latitude  *float64  `json:"latitude,omitempty"`
	Longitude *float64  `json:"longitude,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Связи
	Employees []WarehouseEmployee `gorm:"foreignKey:WarehouseID" json:"-"`
	Products  []Product           `gorm:"foreignKey:WarehouseID" json:"-"`
}

type WarehouseEmployee struct {
	ID          uuid.UUID    `json:"id"`
	WarehouseID uuid.UUID    `json:"warehouse_id"`
	UserID      uint         `json:"user_id"`
	Role        EmployeeRole `json:"role"`
	HiredAt     *time.Time   `json:"hired_at,omitempty"`
	FiredAt     *time.Time   `json:"fired_at,omitempty"`
	IsActive    bool         `json:"is_active"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`

	// Связи
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID" json:"-"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
}
