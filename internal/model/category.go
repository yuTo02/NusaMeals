package model

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Menu      []Menu    `json:"menu"`
}
