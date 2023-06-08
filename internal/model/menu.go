package model

import (
	"time"
)

type Menu struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CategoryID   uint      `json:"category_id" form:"category_id"`
	Name         string    `json:"name" form:"name" gorm:"not null"`
	Price        float64   `json:"price" form:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CategoryMenu string    `json:"category_menu" form:"category_menu" gorm:"type:enum('food','drink')"`
	Calories     string    `json:"calories" form:"calories"`
	City         string    `json:"city" form:"city"`
	Description  string    `json:"description" form:"description"`
	Ingredient   string    `json:"ingredient" form:"ingredient"`
	Images       string    `json:"images" form:"images"`
	Category     Category  `json:"category"`
}

//type CategoryMenu string

//const (
//Food  CategoryMenu = "food"
//Drink CategoryMenu = "drink"
//)
