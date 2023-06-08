package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	MenuUseCase usecase.MenuUseCase
}

func NewMenuController(mu usecase.MenuUseCase) *MenuController {
	return &MenuController{
		MenuUseCase: mu,
	}
}

func (h *MenuController) CreateMenuController(c echo.Context) error {
	var requestMenu request.Menu
	if err := c.Bind(&requestMenu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := c.Validate(requestMenu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err := h.MenuUseCase.CreateMenu(requestMenu)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Menu has been created",
	})
}

func (h *MenuController) GetMenuByIDController(c echo.Context) error {
	ID := c.Param("id")
	menuID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid menu ID",
		})
	}

	menu, err := h.MenuUseCase.GetMenuByID(uint(menuID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    menu,
	})
}

func (h *MenuController) GetAllMenusController(c echo.Context) error {
	menus, err := h.MenuUseCase.GetAllMenus()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    menus,
	})
}

func (h *MenuController) GetMenuByNameController(c echo.Context) error {
	name := c.Param("name")

	menus, err := h.MenuUseCase.GetMenuByName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    menus,
	})
}

func (h *MenuController) GetMenuByCategoryController(c echo.Context) error {
	category := c.Param("category")

	menus, err := h.MenuUseCase.GetMenuByCategory(category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    menus,
	})
}

func (h *MenuController) UpdateMenuController(c echo.Context) error {
	ID := c.Param("id")
	menuID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid menu ID",
		})
	}

	var requestMenu request.Menu
	if err := c.Bind(&requestMenu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := c.Validate(requestMenu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = h.MenuUseCase.UpdateMenu(uint(menuID), requestMenu)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Menu has been updated",
	})
}

func (h *MenuController) DeleteMenuByIDController(c echo.Context) error {
	ID := c.Param("id")
	menuID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid menu ID",
		})
	}

	err = h.MenuUseCase.DeleteMenuByID(uint(menuID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Menu has been deleted",
	})
}
