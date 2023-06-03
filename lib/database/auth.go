package database

import (
	"reglog/config"

	"reglog/models"
)

func Register(user *models.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
