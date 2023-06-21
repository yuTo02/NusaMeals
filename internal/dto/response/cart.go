package response

type CartItemResponse struct {
	ItemID   uint    `json:"item_id"`
	MenuID   uint    `json:"menu_id"`
	MenuName string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
}

type GetCartItemsResponse struct {
	CartItems []CartItemResponse `json:"cart_items"`
}
type GetCartTotalResponse struct {
	Total float64 `json:"total"`
}

type AddCartItemResponse struct {
	Message string `json:"message"`
}

type UpdateCartItemQuantityResponse struct {
	Message string `json:"message"`
}

type RemoveCartItemResponse struct {
	Message string `json:"message"`
}

type ClearCartResponse struct {
	Message string `json:"message"`
}
