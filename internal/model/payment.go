package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	OrderID     uint           `json:"order_id" gorm:"column:order_id"`
	UserID      uint           `json:"user_id" gorm:"column:user_id"`
	Amount      float64        `json:"amount" gorm:"column:amount"`
	Status      string         `json:"status" gorm:"column:status"`
	Method      string         `json:"method" gorm:"column:method"`
	PaymentType string         `json:"payment_type" gorm:"column:payment_type"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
