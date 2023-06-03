package database

import (
	"reglog/config"
	"reglog/models"
)

func GetUserByEmailOrUsername(id string) (user models.User, err error) {
	if err := config.DB.Where("username = ?", id).First(&user).Error; err != nil {
		if err := config.DB.Where("email = ?", id).First(&user).Error; err != nil {
			return models.User{}, err
		}
	}
	return
}

func GetUserById(id int) (user models.User, err error) {
	if err := config.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return
}

func GetProfil(id int) (user models.User, err error) {
	if err = config.DB.First(&user, id).Error; err != nil {
		return
	}
	return
}

func UpdateProfil(req *models.User, username string) error {
	if err := config.DB.Model(&req).Where("username = ?", username).Updates(models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	var user models.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
