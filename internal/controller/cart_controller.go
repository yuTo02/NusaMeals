package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	cartUseCase usecase.CartUseCase
}

func NewCartController(cartUseCase usecase.CartUseCase) *CartController {
	return &CartController{
		cartUseCase: cartUseCase,
	}
}

func (cc *CartController) GetCartTotal(c echo.Context) error {
	cartIDStr := c.Param("cartID")

	// Konversi cartIDStr menjadi uint
	cartID, err := strconv.ParseUint(cartIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid cart ID")
	}

	// Panggil use case untuk mendapatkan total harga cart
	res, err := cc.cartUseCase.GetCartTotal(uint(cartID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get cart total")
	}

	return c.JSON(http.StatusOK, res)
}

func (cc *CartController) AddItemToCart(c echo.Context) error {
	request := new(request.AddCartItemRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	response, err := cc.cartUseCase.AddItemToCart(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to add item to cart")
	}

	return c.JSON(http.StatusOK, response)
}

func (cc *CartController) UpdateCartItemQuantity(c echo.Context) error {
	request := new(request.UpdateCartItemQuantityRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	response, err := cc.cartUseCase.UpdateCartItemQuantity(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update cart item quantity")
	}

	return c.JSON(http.StatusOK, response)
}

func (cc *CartController) GetCartItems(c echo.Context) error {
	cartIDStr := c.Param("cartID")

	// Konversi cartIDStr menjadi uint
	cartID, err := strconv.ParseUint(cartIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid cart ID")
	}

	// Panggil use case untuk mendapatkan daftar item dalam cart
	res, err := cc.cartUseCase.GetCartItems(uint(cartID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get cart items")
	}

	return c.JSON(http.StatusOK, res)
}

func (cc *CartController) RemoveItemFromCart(c echo.Context) error {
	request := new(request.RemoveCartItemRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	response, err := cc.cartUseCase.RemoveItemFromCart(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to remove item from cart")
	}

	return c.JSON(http.StatusOK, response)
}

func (cc *CartController) ClearCart(c echo.Context) error {
	request := new(request.ClearCartRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	response, err := cc.cartUseCase.ClearCart(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to clear cart")
	}

	return c.JSON(http.StatusOK, response)
}
