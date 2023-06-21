package controller

import (
	"net/http"
	"reglog/internal/dto/request"
	"reglog/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
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
	orderDTO := new(request.Order)
	if err := c.Bind(orderDTO); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	createdOrder, err := oc.orderUseCase.CreateOrder(orderDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create order")
	}

	return c.JSON(http.StatusOK, createdOrder)
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

func (oc *OrderController) UpdateOrder(c echo.Context) error {
	orderID := c.Param("orderID")
	// Convert orderID string to uint
	orderIDUint, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	orderDTO := new(request.OrderUpdate)
	if err := c.Bind(orderDTO); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	// Set orderID in the DTO
	orderDTO.ID = uint(orderIDUint)

	updatedOrder, err := oc.orderUseCase.UpdateOrder(orderDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update order")
	}

	return c.JSON(http.StatusOK, updatedOrder)
}

func (oc *OrderController) DeleteOrder(c echo.Context) error {
	orderID := c.Param("orderID")
	// Convert orderID string to uint
	orderIDUint, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	err = oc.orderUseCase.DeleteOrder(uint(orderIDUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete order")
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "the Order deleted successfully"})
}

func (oc *OrderController) GetAllOrders(c echo.Context) error {
	orders, err := oc.orderUseCase.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get orders")
	}

	return c.JSON(http.StatusOK, orders)
}
