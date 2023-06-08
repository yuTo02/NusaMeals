package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type MenuRepository interface {
	CreateMenu(data model.Menu) error
	GetAllMenus() ([]model.Menu, error)
	GetMenuByID(ID uint) (model.Menu, error)
	GetMenuByName(name string) (model.Menu, error)
	GetMenuByCategory(category string) ([]model.Menu, error)
	UpdateMenu(ID uint, data model.Menu) error
	DeleteMenuByID(ID uint) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) CreateMenu(data model.Menu) error {
	if err := r.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *menuRepository) GetAllMenus() ([]model.Menu, error) {
	var menus []model.Menu

	if err := r.db.Model(&model.Menu{}).Find(&menus).Error; err != nil {
		return menus, err
	}

	return menus, nil
}

func (r *menuRepository) GetMenuByID(ID uint) (model.Menu, error) {
	var menu model.Menu

	if err := r.db.Model(&menu).Where("id = ?", ID).Find(&menu).Error; err != nil {
		return menu, err
	}

	return menu, nil
}

func (r *menuRepository) GetMenuByName(name string) (model.Menu, error) {
	var menu model.Menu

	if err := r.db.Model(&model.Menu{}).Where("name = ?", name).First(&menu).Error; err != nil {
		return model.Menu{}, err
	}

	return menu, nil
}

func (r *menuRepository) GetMenuByCategory(category string) ([]model.Menu, error) {
	var menus []model.Menu

	if err := r.db.Model(&model.Menu{}).Where("category = ?", category).Find(&menus).Error; err != nil {
		return menus, err
	}

	return menus, nil
}

func (r *menuRepository) UpdateMenu(ID uint, data model.Menu) error {
	var menu model.Menu

	if err := r.db.Model(&menu).Where("id = ?", ID).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *menuRepository) DeleteMenuByID(ID uint) error {
	var menu model.Menu

	if err := r.db.Where("id = ", ID).Find(&menu).Unscoped().Delete(&menu).Error; err != nil {
		return err
	}

	return nil
}
