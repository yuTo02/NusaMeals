package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint           `gorm:"primaryKey" json:"order_id"`
	UserID      uint           `json:"user_id" gorm:"column:user_id"`
	MenuID      uint           `json:"menu_id" gorm:"column:menu_id"`
	TypeOrder   string         `json:"type_order" gorm:"column:type_order"`
	Quantity    int            `json:"quantity" gorm:"column:quantity"`
	TotalPrice  int            `json:"total_price" gorm:"column:total_price"`
	OrderStatus string         `json:"order_status" gorm:"column:order_status"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
