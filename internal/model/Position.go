package model

type Position struct {
	ID    int     `gorm:"primary_key" json:"id" auto_increment:"true"`
	Name  string  `json:"name" form:"name"`
	Table []Table `json:"table" gorm:"foreignkey:PositionID"`
}
