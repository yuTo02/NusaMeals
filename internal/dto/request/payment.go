package request

type Payment struct {
	ID          uint    `json:"id"`
	OrderID     uint    `json:"order_id"`
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	Method      string  `json:"method"`
	PaymentType string  `json:"payment_type"`
}

type PaymentUpdate struct {
	ID          uint    `json:"id"`
	OrderID     uint    `json:"order_id"`
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	Method      string  `json:"method"`
	PaymentType string  `json:"payment_type"`
}
