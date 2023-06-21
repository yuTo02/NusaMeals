package response

type GetMenuResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Calories    string  `json:"calories"`
	City        string  `json:"city"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Images      string  `json:"images"`
	CategoryID  uint    `json:"category_id"`
	Category    string  `json:"category"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetCategoryResponse struct {
	ID    uint              `json:"id"`
	Name  string            `json:"name"`
	Menus []GetMenuResponse `json:"menus"`
}
