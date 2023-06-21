package model

import "github.com/jinzhu/gorm"

type CartItem struct {
	gorm.Model
	CartID   uint    `gorm:"index" json:"cart_id"`
	MenuID   uint    `gorm:"index" json:"menu_id"`
	Menu     Menu    `json:"menu" gorm:"constraint:OnDelete:CASCADE"`
	Quantity float64 `json:"quantity"`
	Subtotal float64 `json:"subtotal"`
	Cart     Cart    `gorm:"foreignkey:CartID"`
}
