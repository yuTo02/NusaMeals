package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryController(cu usecase.CategoryUseCase) *CategoryController {
	return &CategoryController{
		CategoryUseCase: cu,
	}
}

func (cc *CategoryController) GetCategoryController(c echo.Context) error {
	categories, err := cc.CategoryUseCase.GetCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all categories",
		"data":    categories,
	})
}

func (cc *CategoryController) CreateCategoryController(c echo.Context) error {
	var req request.CreateCategory
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := cc.CategoryUseCase.CreateCategory(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create category",
	})
}
