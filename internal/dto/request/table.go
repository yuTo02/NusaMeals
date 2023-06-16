package request

type Table struct {
	Number_Table uint   `json:"number_table" form:"number_table"`
	Seat         uint   `json:"seat" form:"seat"`
	Position     string `json:"position" form:"position"`
	Status       string `json:"status" form:"status"`
	Location     string `json:"location" form:"location"`
	Images       string `json:"images" form:"images"`
}
