package request

type Product struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Stock uint   `json:"stock" form:"stock" validate:"required"`
	Type  string `json:"type" form:"type" validate:"required"`
}
