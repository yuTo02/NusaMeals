package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type CategoryRepository interface {
	CreateCategory(data model.Category) error
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(ID uint) (model.Category, error)
	GetMenuByCategory(categoryID uint) ([]model.Menu, error)
	GetCategoryByName(name string) (model.Category, error)
	UpdateCategory(data model.Category) error
	DeleteCategory(ID uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(data model.Category) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) GetCategoryByID(ID uint) (model.Category, error) {
	var category model.Category
	err := r.db.First(&category, ID).Error
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) GetMenuByCategory(categoryID uint) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Where("category_id = ?", categoryID).Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *categoryRepository) GetCategoryByName(name string) (model.Category, error) {
	var category model.Category
	err := r.db.Preload("Menus").Where("name = ?", name).First(&category).Error
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) UpdateCategory(data model.Category) error {
	err := r.db.Save(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) DeleteCategory(ID uint) error {
	err := r.db.Delete(&model.Category{}, ID).Error
	if err != nil {
		return err
	}
	return nil
}
