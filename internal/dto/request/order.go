package request

import "time"

type Order struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	MenuID      uint      `json:"menu_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  int       `json:"total_price"`
	TypeOrder   string    `json:"type_order"`
	OrderStatus string    `json:"order_status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type OrderUpdate struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	MenuID      uint   `json:"menu_id"`
	TypeOrder   string `json:"type_order"`
	OrderStatus string `json:"order_status"`
	TotalPrice  int    `json:"total_price"`
}
