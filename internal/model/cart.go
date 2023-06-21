package model

import (
	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint       `gorm:"unique_index" json:"user_id"`
	Items  []CartItem `json:"items" gorm:"foreignkey:CartID"`
}
