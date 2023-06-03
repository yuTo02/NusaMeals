package database

import (
	"reglog/internal/common/config"
	"reglog/internal/model"
)

func Register(user *model.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
