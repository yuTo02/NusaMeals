package model

import (
	"gorm.io/gorm"
	"time"
)

type Table struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `json:"name" gorm:"column:name"`
	Capacity   int            `json:"capacity" gorm:"column:capacity"`
	IsOccupied bool           `json:"isOccupied" gorm:"column:is_occupied"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
