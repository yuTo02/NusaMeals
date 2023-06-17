package repository

import (
	"github.com/jinzhu/gorm"
	"reglog/internal/model"
)

type MenuItemRepository interface {
	CreateMenuItem(menuItem *model.MenuItem) (*model.MenuItem, error)
	UpdateMenuItem(menuItem *model.MenuItem) (*model.MenuItem, error)
	DeleteMenuItem(menuItemID uint) error
	GetMenuItemByID(menuItemID uint) (*model.MenuItem, error)
	GetAllMenuItems() ([]model.MenuItem, error)
}

type menuItemRepository struct {
	db *gorm.DB
}

func NewMenuItemRepository(db *gorm.DB) *menuItemRepository {
	return &menuItemRepository{
		db: db,
	}
}

func (r *menuItemRepository) CreateMenuItem(menuItem *model.MenuItem) (*model.MenuItem, error) {
	result := r.db.Create(menuItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return menuItem, nil
}

func (r *menuItemRepository) UpdateMenuItem(menuItem *model.MenuItem) (*model.MenuItem, error) {
	result := r.db.Save(menuItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return menuItem, nil
}

func (r *menuItemRepository) DeleteMenuItem(menuItemID uint) error {
	return r.db.Delete(&model.MenuItem{}, menuItemID).Error
}

func (r *menuItemRepository) GetMenuItemByID(menuItemID uint) (*model.MenuItem, error) {
	var menuItem model.MenuItem
	err := r.db.First(&menuItem, menuItemID).Error
	if err != nil {
		return nil, err
	}
	return &menuItem, nil
}

func (r *menuItemRepository) GetAllMenuItems() ([]model.MenuItem, error) {
	var menuItems []model.MenuItem
	err := r.db.Find(&menuItems).Error
	if err != nil {
		return nil, err
	}
	return menuItems, nil
}
