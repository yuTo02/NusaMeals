package response

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name" form:"name"`
	Stock uint   `json:"stock" form:"stock"`
	Type  string `json:"type" form:"type"`
}
