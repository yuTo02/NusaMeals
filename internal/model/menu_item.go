package model

import (
	"gorm.io/gorm"
	"time"
)

type MenuItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name" gorm:"column:name"`
	Price     float64        `json:"price" gorm:"column:price"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
