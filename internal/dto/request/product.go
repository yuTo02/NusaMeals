package request

type Product struct {
	Name        string  `json:"name" form:"name" validate:"required"`
	Stock       uint    `json:"stock" form:"stock" validate:"required"`
	Category    string  `json:"category" form:"category" validate:"required"`
	Price       float64 `json:"price" form:"price" validate:"required"`
	Calories    uint    `json:"calories" form:"calories" validate:"required"`
	City        string  `json:"city" form:"city" validate:"required"`
	Description string  `json:"description" form:"description" validate:"required"`
	Ingredients string  `json:"ingredients" form:"ingredients" validate:"required"`
}
