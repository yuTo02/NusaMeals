package response

import "time"

type Order struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	MenuID      uint      `json:"menu_id"`
	Quantity    int       `json:"quantity"`
	TypeOrder   string    `json:"type_order"`
	TotalPrice  float64   `json:"total_price"`
	OrderStatus string    `json:"order_status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type OrderUpdate struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	MenuID      uint    `json:"menu_id"`
	TypeOrder   string  `json:"type_order"`
	OrderStatus string  `json:"order_status"`
	TotalPrice  float64 `json:"total_price"`
}
