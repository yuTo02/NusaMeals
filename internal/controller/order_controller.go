package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reglog/internal/dtos"
	"reglog/internal/usecase"
	"strconv"
)

type OrderController struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderController(orderUseCase usecase.OrderUseCase) *OrderController {
	return &OrderController{
		orderUseCase: orderUseCase,
	}
}

func (oc *OrderController) CreateOrder(c echo.Context) error {
	orderDTO := new(dtos.OrderDTO)
	if err := c.Bind(orderDTO); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	createdOrder, err := oc.orderUseCase.CreateOrder(orderDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create order")
	}

	return c.JSON(http.StatusCreated, createdOrder)
}

func (oc *OrderController) GetOrderByID(c echo.Context) error {
	orderID := c.Param("orderID")
	// Convert orderID string to uint
	orderIDUint, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	order, err := oc.orderUseCase.GetOrderByID(uint(orderIDUint))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Order not found")
	}

	return c.JSON(http.StatusOK, order)
}
