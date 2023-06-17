package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reglog/internal/dtos"
	"reglog/internal/usecase"
	"strconv"
)

type OrderItemController struct {
	orderItemUseCase usecase.OrderItemUseCase
}

func NewOrderItemController(orderItemUseCase usecase.OrderItemUseCase) *OrderItemController {
	return &OrderItemController{
		orderItemUseCase: orderItemUseCase,
	}
}

func (c *OrderItemController) AddOrderItem(ctx echo.Context) error {
	orderID, err := strconv.ParseUint(ctx.Param("order_id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	var orderItemDTO dtos.OrderItemDTO
	if err := ctx.Bind(&orderItemDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid order item data")
	}

	orderItemDTO.OrderID = uint(orderID)

	createdOrderItem, err := c.orderItemUseCase.AddOrderItem(&orderItemDTO)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to add order item")
	}

	return ctx.JSON(http.StatusOK, createdOrderItem)
}

func (c *OrderItemController) UpdateOrderItem(ctx echo.Context) error {
	orderItemID, err := strconv.ParseUint(ctx.Param("order_item_id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid order item ID")
	}

	var orderItemDTO dtos.OrderItemDTO
	if err := ctx.Bind(&orderItemDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid order item data")
	}

	updatedOrderItem, err := c.orderItemUseCase.UpdateOrderItem(uint(orderItemID), &orderItemDTO)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update order item")
	}

	return ctx.JSON(http.StatusOK, updatedOrderItem)
}

func (c *OrderItemController) RemoveOrderItem(ctx echo.Context) error {
	orderItemID, err := strconv.ParseUint(ctx.Param("order_item_id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid order item ID")
	}

	err = c.orderItemUseCase.RemoveOrderItem(uint(orderItemID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to remove order item")
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *OrderItemController) GetOrderItemsByOrderID(ctx echo.Context) error {
	orderID, err := strconv.ParseUint(ctx.Param("order_id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	orderItems, err := c.orderItemUseCase.GetOrderItemsByOrderID(uint(orderID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get order items")
	}

	return ctx.JSON(http.StatusOK, orderItems)
}
