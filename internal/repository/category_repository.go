package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type CategoryRepository interface {
	CreateCategory(data model.Category) error
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(ID uint) (model.Category, error)
	// Implement other operations as needed
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
