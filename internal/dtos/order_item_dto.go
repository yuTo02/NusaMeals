package dtos

type OrderItemDTO struct {
	ID         uint    `json:"id"`
	OrderID    uint    `json:"order_id"`
	MenuItemID uint    `json:"menu_item_id"`
	Quantity   int     `json:"quantity"`
	Subtotal   float64 `json:"subtotal"`
}
