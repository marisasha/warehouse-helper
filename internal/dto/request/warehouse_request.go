package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/marisasha/warehouse-helper/internal/models"
)

type WarehouseEmployee struct {
	WarehouseID uuid.UUID           `json:"-"`
	UserID      uint                `json:"user_id"`
	Role        models.EmployeeRole `json:"role"`
	HiredAt     *time.Time          `json:"hired_at,omitempty"`
}
