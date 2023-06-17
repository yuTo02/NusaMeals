package dtos

type TableDTO struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Capacity   int    `json:"capacity"`
	IsOccupied bool   `json:"isOccupied"`
}
