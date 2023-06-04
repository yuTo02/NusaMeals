package model

import "time"

type Product struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name" form:"name" gorm:"not null"`
	Stock       uint      `json:"stock" form:"stock" gorm:"not null"`
	Category    string    `json:"category" form:"category" gorm:"type:enum('makanan','minuman')"`
	Price       float64   `json:"price" form:"price" gorm:"not null"`
	Calories    uint      `json:"calories" form:"calories" gorm:"null"`
	City        string    `json:"city" form:"city" gorm:"null"`
	Description string    `json:"description" form:"description" gorm:"null"`
	Ingredients string    `json:"ingredients" form:"ingredients" gorm:"null"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
