package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reglog/internal/dtos"
	"reglog/internal/usecase"
	"strconv"
)

type MenuItemController struct {
	menuItemUseCase usecase.MenuItemUseCase
}

func NewMenuItemController(menuItemUseCase usecase.MenuItemUseCase) *MenuItemController {
	return &MenuItemController{
		menuItemUseCase: menuItemUseCase,
	}
}

func (c *MenuItemController) CreateMenuItem(ctx echo.Context) error {
	var menuItemDTO dtos.MenuItemDTO
	if err := ctx.Bind(&menuItemDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdMenuItem, err := c.menuItemUseCase.CreateMenuItem(&menuItemDTO)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, createdMenuItem)
}

func (c *MenuItemController) UpdateMenuItem(ctx echo.Context) error {
	menuItemIDStr := ctx.Param("id")
	menuItemID, err := strconv.ParseUint(menuItemIDStr, 10, 64)

	var menuItemDTO dtos.MenuItemDTO
	if err := ctx.Bind(&menuItemDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedMenuItem, err := c.menuItemUseCase.UpdateMenuItem(uint(menuItemID), &menuItemDTO)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, updatedMenuItem)
}

func (c *MenuItemController) DeleteMenuItem(ctx echo.Context) error {
	menuItemIDStr := ctx.Param("id")
	menuItemID, err := strconv.ParseUint(menuItemIDStr, 10, 64)

	err = c.menuItemUseCase.DeleteMenuItem(uint(menuItemID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Menu item deleted successfully"})
}

func (c *MenuItemController) GetMenuItemByID(ctx echo.Context) error {
	menuItemIDStr := ctx.Param("id")
	menuItemID, err := strconv.ParseUint(menuItemIDStr, 10, 64)

	menuItem, err := c.menuItemUseCase.GetMenuItemByID(uint(menuItemID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, menuItem)
}

func (c *MenuItemController) GetAllMenuItems(ctx echo.Context) error {
	menuItems, err := c.menuItemUseCase.GetAllMenuItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, menuItems)
}
