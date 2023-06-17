package usecase

import (
	"reglog/internal/dtos"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type MenuItemUseCase interface {
	CreateMenuItem(menuItemDTO *dtos.MenuItemDTO) (*dtos.MenuItemDTO, error)
	UpdateMenuItem(menuItemID uint, menuItemDTO *dtos.MenuItemDTO) (*dtos.MenuItemDTO, error)
	DeleteMenuItem(menuItemID uint) error
	GetMenuItemByID(menuItemID uint) (*dtos.MenuItemDTO, error)
	GetAllMenuItems() ([]dtos.MenuItemDTO, error)
}

type menuItemUseCase struct {
	menuItemRepo repository.MenuItemRepository
}

func NewMenuItemUseCase(menuItemRepo repository.MenuItemRepository) MenuItemUseCase {
	return &menuItemUseCase{
		menuItemRepo: menuItemRepo,
	}
}

func (m *menuItemUseCase) CreateMenuItem(menuItemDTO *dtos.MenuItemDTO) (*dtos.MenuItemDTO, error) {
	menuItem := &model.MenuItem{
		Name:  menuItemDTO.Name,
		Price: menuItemDTO.Price,
		// Set other fields accordingly
	}

	menuItem, err := m.menuItemRepo.CreateMenuItem(menuItem)
	if err != nil {
		return nil, err
	}

	createdMenuItemDTO := &dtos.MenuItemDTO{
		ID:    menuItem.ID,
		Name:  menuItem.Name,
		Price: menuItem.Price,
		// Set other fields accordingly
	}

	return createdMenuItemDTO, nil
}

func (m *menuItemUseCase) UpdateMenuItem(menuItemID uint, menuItemDTO *dtos.MenuItemDTO) (*dtos.MenuItemDTO, error) {
	menuItem, err := m.menuItemRepo.GetMenuItemByID(menuItemID)
	if err != nil {
		return nil, err
	}

	// Update the menu item entity with the new data
	menuItem.Name = menuItemDTO.Name
	menuItem.Price = menuItemDTO.Price
	// Update other fields accordingly

	updatedMenuItem, err := m.menuItemRepo.UpdateMenuItem(menuItem)
	if err != nil {
		return nil, err
	}

	updatedMenuItemDTO := &dtos.MenuItemDTO{
		ID:    updatedMenuItem.ID,
		Name:  updatedMenuItem.Name,
		Price: updatedMenuItem.Price,
		// Set other fields accordingly
	}

	return updatedMenuItemDTO, nil
}

func (m *menuItemUseCase) DeleteMenuItem(menuItemID uint) error {
	return m.menuItemRepo.DeleteMenuItem(menuItemID)
}

func (m *menuItemUseCase) GetMenuItemByID(menuItemID uint) (*dtos.MenuItemDTO, error) {
	menuItem, err := m.menuItemRepo.GetMenuItemByID(menuItemID)
	if err != nil {
		return nil, err
	}

	menuItemDTO := &dtos.MenuItemDTO{
		ID:    menuItem.ID,
		Name:  menuItem.Name,
		Price: menuItem.Price,
		// Set other fields accordingly
	}

	return menuItemDTO, nil
}

func (m *menuItemUseCase) GetAllMenuItems() ([]dtos.MenuItemDTO, error) {
	menuItems, err := m.menuItemRepo.GetAllMenuItems()
	if err != nil {
		return nil, err
	}

	var menuItemDTOs []dtos.MenuItemDTO
	for _, menuItem := range menuItems {
		menuItemDTO := dtos.MenuItemDTO{
			ID:    menuItem.ID,
			Name:  menuItem.Name,
			Price: menuItem.Price,
			// Set other fields accordingly
		}
		menuItemDTOs = append(menuItemDTOs, menuItemDTO)
	}

	return menuItemDTOs, nil
}
