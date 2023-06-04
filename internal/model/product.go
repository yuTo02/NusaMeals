package model

import "time"

type Product struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" form:"name" gorm:"not null"`
	Stock     uint      `json:"stock" form:"stock" gorm:"not null"`
	Type      string    `json:"type" form:"type" gorm:"type:enum('makanan','minuman')"`
	CreatedAt time.Time `json:"created_at"`
}
