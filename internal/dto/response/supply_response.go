package response

import (
	"github.com/google/uuid"
	"github.com/marisasha/warehouse-helper/internal/models"
)

type UpdateSupplyStatus struct {
	NewStatus models.SupplyStatus `json:"new_status"`
	BatchID   uuid.UUID           `json:"batch_id"`
}
