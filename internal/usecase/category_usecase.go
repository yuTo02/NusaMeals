package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type CategoryUseCase interface {
	CreateCategory(req request.CreateCategory) error
	GetCategories() ([]string, error)
}

type categoryUseCase struct {
	CategoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		CategoryRepo: categoryRepo,
	}
}

func (uc *categoryUseCase) CreateCategory(req request.CreateCategory) error {
	category := model.Category{
		Name: req.Name,
	}
	err := uc.CategoryRepo.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (uc *categoryUseCase) GetCategories() ([]string, error) {
	categories, err := uc.CategoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	var response []string
	for _, category := range categories {
		response = append(response, category.Name)
	}

	return response, nil
}
