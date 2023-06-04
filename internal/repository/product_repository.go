package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	CreateProduct(data model.Product) error
	GetAllProducts() ([]model.Product, error)
	GetProductByID(ID string) (model.Product, error)
	GetProductByCategory(category string) (model.Product, error)
	UpdateProduct(ID uint, data model.Product) error
	//DeleteProductByID(ID uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(data model.Product) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product

	if err := r.db.Model(&model.Product{}).Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (r *productRepository) GetProductByID(ID string) (model.Product, error) {

	var product model.Product
	if err := r.db.Model(&product).Where("id = ?", ID).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) GetProductByCategory(category string) (model.Product, error) {

	var product model.Product
	if err := r.db.Model(&product).Where("category = ?", category).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
