package dtos

import "time"

type OrderDTO struct {
	ID          uint           `json:"id"`
	UserID      uint           `json:"user_id"`
	OrderItems  []OrderItemDTO `json:"items"`
	TotalAmount float64        `json:"total_amount"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}
