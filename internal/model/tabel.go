package model

type Tabel struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Number_Table uint   `json:"number_table" form:"number_table" gorm:"not null"`
	Seat         uint   `json:"seat" form:"seat" gorm:"unique;not null"`
	Position     string `json:"position" form:"position" gorm:"type:enum('indoor','outdoor');default:'indoor'"`
	Status       string `json:"status" form:"status" gorm:"type:enum('unavailable','available');default:'available'"`
	Location     string `json:"location" form:"location" gorm:"uniqe;not null"`
	Images       string `json:"images" form:"images"`
}
