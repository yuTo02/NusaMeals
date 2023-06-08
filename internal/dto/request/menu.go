package request

type Menu struct {
	Name         string  `json:"name" form:"name" binding:"required"`
	Price        float64 `json:"price" form:"price" binding:"required"`
	CategoryID   uint    `json:"category_id" form:"category_id" binding:"required"`
	CategoryMenu string  `json:"category_menu" form:"category_menu"`
	Calories     string  `json:"calories" form:"calories"`
	Description  string  `json:"description" form:"description"`
	Ingredient   string  `json:"ingredient" form:"ingredient"`
}
