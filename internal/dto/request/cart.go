package request

type AddCartItemRequest struct {
	CartID   uint    `json:"cart_id"`
	MenuID   uint    `json:"menu_id"`
	Quantity float64 `json:"quantity"`
}

type UpdateCartItemQuantityRequest struct {
	CartItemID uint    `json:"cart_item_id"`
	Quantity   float64 `json:"quantity"`
}

type RemoveCartItemRequest struct {
	CartItemID uint `json:"item_id"`
}

type ClearCartRequest struct {
	CartID uint `json:"cart_id"`
}
