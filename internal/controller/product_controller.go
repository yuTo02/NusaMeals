package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
)

type ProductController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(pu usecase.ProductUseCase) *ProductController {
	return &ProductController{
		ProductUseCase: pu,
	}
}

func (h *ProductController) CreateProductController(c echo.Context) error {
	var requestProduct request.Product
	if err := c.Bind(&requestProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := c.Validate(requestProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err := h.ProductUseCase.CreateProduct(requestProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product has been created",
	})
}
