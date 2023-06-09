package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type MenuUseCase interface {
	CreateMenu(request request.Menu) error
	GetMenuByID(ID uint) (model.Menu, error)
	GetAllMenus() ([]model.Menu, error)
	GetMenusByName(name string) ([]model.Menu, error)
	GetMenusByCategory(categoryTypeID uint) ([]model.Menu, error)
	GetMenusByCategoryName(categoryName string) ([]model.Menu, error)
	UpdateMenu(ID uint, request request.UpdateMenu) error
	DeleteMenuByID(ID uint) error
}

type menuUseCase struct {
	MenuRepo repository.MenuRepository
}

func NewMenuUseCase(menuRepo repository.MenuRepository) MenuUseCase {
	return &menuUseCase{
		MenuRepo: menuRepo,
	}
}

func (uc *menuUseCase) CreateMenu(request request.Menu) error {
	err := uc.MenuRepo.CreateMenu(&request)
	if err != nil {
		return err
	}
	return nil
}

func (uc *menuUseCase) GetMenuByID(ID uint) (model.Menu, error) {
	menu, err := uc.MenuRepo.GetMenuByID(ID)
	if err != nil {
		return model.Menu{}, err
	}
	return menu, nil
}

func (uc *menuUseCase) GetAllMenus() ([]model.Menu, error) {
	menus, err := uc.MenuRepo.GetMenus()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (uc *menuUseCase) GetMenusByName(name string) ([]model.Menu, error) {
	menus, err := uc.MenuRepo.GetMenusByName(name)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (uc *menuUseCase) GetMenusByCategory(categoryID uint) ([]model.Menu, error) {
	menus, err := uc.MenuRepo.GetMenusByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (uc *menuUseCase) GetMenusByCategoryName(categoryName string) ([]model.Menu, error) {
	menus, err := uc.MenuRepo.GetMenusByCategoryName(categoryName)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (uc *menuUseCase) UpdateMenu(ID uint, request request.UpdateMenu) error {
	menu, err := uc.MenuRepo.GetMenuByID(ID)
	if err != nil {
		return err
	}

	menu.Name = request.Name
	menu.Price = request.Price
	menu.Calories = request.Calories
	menu.City = request.City
	menu.Description = request.Description
	menu.Ingredient = request.Ingredient
	menu.Images = request.Images
	menu.CategoryID = request.CategoryID

	err = uc.MenuRepo.UpdateMenu(&menu)
	if err != nil {
		return err
	}
	return nil
}

func (uc *menuUseCase) DeleteMenuByID(ID uint) error {
	err := uc.MenuRepo.DeleteMenuByID(ID)
	if err != nil {
		return err
	}
	return nil
}
