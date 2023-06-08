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
	GetMenuByName(name string) (model.Menu, error)
	GetMenuByCategory(category string) ([]model.Menu, error)
	UpdateMenu(ID uint, request request.Menu) error
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

func (u *menuUseCase) CreateMenu(request request.Menu) error {
	menu := model.Menu{
		Name:         request.Name,
		Price:        request.Price,
		CategoryID:   request.CategoryID,
		CategoryMenu: request.CategoryMenu,
		Calories:     request.Calories,
		Description:  request.Description,
		Ingredient:   request.Ingredient,
	}

	err := u.MenuRepo.CreateMenu(menu)
	if err != nil {
		return err
	}

	return nil
}

func (u *menuUseCase) GetMenuByID(ID uint) (model.Menu, error) {
	menu, err := u.MenuRepo.GetMenuByID(ID)
	if err != nil {
		return model.Menu{}, err
	}

	return menu, nil
}

func (u *menuUseCase) GetAllMenus() ([]model.Menu, error) {
	menus, err := u.MenuRepo.GetAllMenus()
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (u *menuUseCase) GetMenuByName(name string) (model.Menu, error) {
	menu, err := u.MenuRepo.GetMenuByName(name)
	if err != nil {
		return model.Menu{}, err
	}

	return menu, nil
}

func (u *menuUseCase) GetMenuByCategory(category string) ([]model.Menu, error) {
	menus, err := u.MenuRepo.GetMenuByCategory(category)
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (u *menuUseCase) UpdateMenu(ID uint, request request.Menu) error {
	menu, err := u.MenuRepo.GetMenuByID(ID)
	if err != nil {
		return err
	}

	menu.Name = request.Name
	menu.Price = request.Price
	menu.CategoryID = request.CategoryID
	menu.CategoryMenu = request.CategoryMenu
	menu.Calories = request.Calories
	menu.Description = request.Description
	menu.Ingredient = request.Ingredient

	err = u.MenuRepo.UpdateMenu(ID, menu)
	if err != nil {
		return err
	}

	return nil
}

func (u *menuUseCase) DeleteMenuByID(ID uint) error {
	err := u.MenuRepo.DeleteMenuByID(ID)
	if err != nil {
		return err
	}

	return nil
}
