package model

type Table struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Number_Table uint   `json:"number_table" form:"number_table" gorm:"not null"`
	Seat         uint   `json:"seat" form:"seat" gorm:"unique;not null"`
	PositionID   uint   `json:"position_id" form:"position_id"`
	Status       string `json:"status" form:"status" gorm:"type:enum('unavailable','available');default:'available'"`
	Location     string `json:"location" form:"location" gorm:"uniqe;not null"`
	Images       string `json:"images" form:"images"`
}
