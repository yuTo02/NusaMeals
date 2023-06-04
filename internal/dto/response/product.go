package response

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" form:"name"`
	Stock       uint    `json:"stock" form:"stock"`
	Type        string  `json:"type" form:"type"`
	Category    string  `json:"category" form:"category"`
	Price       float64 `json:"price" form:"price"`
	Calories    uint    `json:"calories" form:"calories"`
	City        string  `json:"city" form:"city" gorm:"null"`
	Description string  `json:"description" form:"description"`
	Ingredients string  `json:"ingredients" form:"ingredients"`
}
