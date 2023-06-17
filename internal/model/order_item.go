package model

import (
	"gorm.io/gorm"
	"time"
)

type OrderItem struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	OrderID    uint           `json:"order_id" gorm:"column:order_id"`
	MenuItemID uint           `json:"menu_item_id" gorm:"column:menu_item_id"`
	Quantity   int            `json:"quantity" gorm:"column:quantity"`
	Subtotal   float64        `json:"sub_total" gorm:"column:sub_total"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
