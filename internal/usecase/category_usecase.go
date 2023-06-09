package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type CategoryUseCase interface {
	CreateCategory(req request.CreateCategory) error
	GetCategories() ([]string, error)
	GetCategoryByID(ID uint) (model.Category, error)
	GetMenuByCategory(categoryID uint) ([]model.Menu, error)
	GetCategoryByName(name string) (model.Category, error)
	UpdateCategory(ID uint, req request.UpdateCategory) error
	DeleteCategory(ID uint) error
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

func (uc *categoryUseCase) GetCategoryByID(ID uint) (model.Category, error) {
	category, err := uc.CategoryRepo.GetCategoryByID(ID)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (uc *categoryUseCase) GetMenuByCategory(categoryID uint) ([]model.Menu, error) {
	menus, err := uc.CategoryRepo.GetMenuByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (uc *categoryUseCase) GetCategoryByName(name string) (model.Category, error) {
	category, err := uc.CategoryRepo.GetCategoryByName(name)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (uc *categoryUseCase) UpdateCategory(ID uint, req request.UpdateCategory) error {
	category, err := uc.CategoryRepo.GetCategoryByID(ID)
	if err != nil {
		return err
	}

	// Update the category fields based on the request
	category.Name = req.Name

	err = uc.CategoryRepo.UpdateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (uc *categoryUseCase) DeleteCategory(ID uint) error {
	err := uc.CategoryRepo.DeleteCategory(ID)
	if err != nil {
		return err
	}
	return nil
}
