package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `json:"user_id" gorm:"column:user_id"`
	OrderItems  []OrderItem    `json:"order_items" gorm:"foreignKey:CartID"`
	TotalAmount float64        `json:"total_amount" gorm:"column:total_amount"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
