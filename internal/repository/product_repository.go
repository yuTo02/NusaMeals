package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	CreateProduct(data model.Product) error
	//GetAllProducts() ([]model.Product, error)
	//GetProductByID(ID string) (model.Product, error)
	//GetProductByType(email string) (model.Product, error)
	//UpdateProduct(ID uint, data model.Product) error
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
