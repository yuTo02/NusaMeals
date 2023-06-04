package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type ProductUseCase interface {
	CreateProduct(request request.Product) error
}

type productUseCase struct {
	ProductRepo repository.ProductRepository
}

func NewProductUseCase(pr repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		ProductRepo: pr,
	}
}

func (p productUseCase) CreateProduct(request request.Product) error {
	product := model.Product{
		Name:  request.Name,
		Stock: request.Stock,
		Type:  request.Type,
	}

	err := p.ProductRepo.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
