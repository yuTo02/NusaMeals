package response

type GetMenuResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"category_id"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	// tambahkan field lain yang diperlukan
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
