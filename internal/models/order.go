package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	OrderPending    OrderStatus = "pending"    // ожидает обработки
	OrderProcessing OrderStatus = "processing" // ожидает исполнения
	OrderReady      OrderStatus = "ready"      // готова к отгрузке (опционально)
	OrderCompleted  OrderStatus = "completed"  // успешно
	OrderRejected   OrderStatus = "rejected"   // отказано
	OrderCancelled  OrderStatus = "cancelled"  // отменено заказчиком
)

type Order struct {
	ID                uuid.UUID       `json:"id"`
	WarehouseID       uuid.UUID       `json:"warehouse_id"`
	CustomerID        uint64          `json:"customer_id"`
	AssigneeID        *uint64         `json:"assignee_id"`
	OrderNumber       string          `json:"order_number"`
	Status            OrderStatus     `json:"status"`
	RequestedShipDate *time.Time      `json:"requested_ship_date,omitempty"`
	ShipWindowStart   *time.Time      `json:"ship_window_start,omitempty"`
	ShipWindowEnd     *time.Time      `json:"ship_window_end,omitempty"`
	CustomerComment   *string         `json:"customer_comment,omitempty"`
	StaffComment      *string         `json:"staff_comment,omitempty"`
	RejectionReason   *string         `json:"rejection_reason,omitempty"`
	TotalAmount       decimal.Decimal `json:"total_amount"`
	ConfirmedAt       *time.Time      `json:"confirmed_at,omitempty"`
	ShippedAt         *time.Time      `json:"shipped_at,omitempty"`
	CompletedAt       *time.Time      `json:"completed_at,omitempty"`
	CancelledAt       *time.Time      `json:"cancelled_at,omitempty"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`

	// Связи
	Warehouse Warehouse            `gorm:"foreignKey:WarehouseID" json:"-"`
	Customer  User                 `gorm:"foreignKey:CustomerID" json:"customer"`
	Assignee  *User                `gorm:"foreignKey:AssigneeID" json:"assignee"`
	Items     []OrderItem          `gorm:"foreignKey:OrderID" json:"items"`
	History   []OrderStatusHistory `gorm:"foreignKey:OrderID" json:"-"`
	Messages  []OrderMessage       `gorm:"foreignKey:OrderID" json:"-"`
}

type OrderItem struct {
	ID           uuid.UUID       `json:"id"`
	OrderID      uuid.UUID       `json:"order_id"`
	ProductID    uuid.UUID       `json:"product_id"`
	Quantity     float64         `json:"quantity"`
	PricePerUnit decimal.Decimal `json:"price_per_unit"`
	Total        decimal.Decimal `json:"total"`
	Comment      *string         `json:"comment,omitempty"`

	Order   Order   `gorm:"foreignKey:OrderID" json:"-"`
	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}

type OrderItemBatchAllocation struct {
	ID          uuid.UUID `json:"id"`
	OrderItemID uuid.UUID `json:"order_item_id"`
	BatchID     uuid.UUID `json:"batch_id"`
	Quantity    float64   `json:"quantity"`

	OrderItem OrderItem `gorm:"foreignKey:OrderItemID" json:"-"`
	Batch     Batch     `gorm:"foreignKey:BatchID" json:"-"`
}

type OrderStatusHistory struct {
	ID        uuid.UUID   `json:"id"`
	OrderID   uuid.UUID   `json:"order_id"`
	OldStatus *string     `json:"old_status,omitempty"`
	NewStatus OrderStatus `json:"new_status"`
	ChangedBy *uint64     `json:"changed_by,omitempty"`
	Comment   *string     `json:"comment,omitempty"`
	CreatedAt time.Time   `json:"created_at"`

	Order Order `gorm:"foreignKey:OrderID" json:"-"`
}

type OrderMessage struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	SenderID  uint64    `json:"sender_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`

	Order  Order `gorm:"foreignKey:OrderID" json:"-"`
	Sender User  `gorm:"foreignKey:SenderID" json:"sender"`
}
