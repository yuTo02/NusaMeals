package database

import (
	"reglog/internal/common/config"
	"reglog/internal/model"
)

func GetUserByEmailOrUsername(id string) (user model.User, err error) {
	if err := config.DB.Where("username = ?", id).First(&user).Error; err != nil {
		if err := config.DB.Where("email = ?", id).First(&user).Error; err != nil {
			return model.User{}, err
		}
	}
	return
}

func GetUserById(id int) (user model.User, err error) {
	if err := config.DB.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return
}

func GetProfil(id int) (user model.User, err error) {
	if err = config.DB.First(&user, id).Error; err != nil {
		return
	}
	return
}

func UpdateProfil(req *model.User, username string) error {
	if err := config.DB.Model(&req).Where("username = ?", username).Updates(model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	var user model.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
