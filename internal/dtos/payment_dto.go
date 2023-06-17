package dtos

type PaymentDTO struct {
	ID          uint    `json:"id"`
	OrderID     uint    `json:"order_id"`
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Method      string  `json:"method"`
	PaymentType string  `json:"payment_type"`
}
