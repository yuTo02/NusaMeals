package usecase

import (
	"reglog/internal/dtos"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type OrderItemUseCase interface {
	AddOrderItem(orderItemDTO *dtos.OrderItemDTO) (*dtos.OrderItemDTO, error)
	UpdateOrderItem(orderItemID uint, orderItemDTO *dtos.OrderItemDTO) (*dtos.OrderItemDTO, error)
	RemoveOrderItem(orderItemID uint) error
	GetOrderItemsByOrderID(orderID uint) ([]dtos.OrderItemDTO, error)
}

type orderItemUseCase struct {
	orderItemRepo repository.OrderItemRepository
}

func NewOrderItemUseCase(orderItemRepo repository.OrderItemRepository) OrderItemUseCase {
	return &orderItemUseCase{
		orderItemRepo: orderItemRepo,
	}
}

func (c *orderItemUseCase) AddOrderItem(orderItemDTO *dtos.OrderItemDTO) (*dtos.OrderItemDTO, error) {
	orderItem := &model.OrderItem{
		OrderID:    orderItemDTO.OrderID,
		MenuItemID: orderItemDTO.MenuItemID,
		Quantity:   orderItemDTO.Quantity,
		Subtotal:   orderItemDTO.Subtotal,
		// Set other fields accordingly
	}

	createdOrderItem, err := c.orderItemRepo.CreateOrderItem(orderItem)
	if err != nil {
		return nil, err
	}

	createdOrderItemDTO := &dtos.OrderItemDTO{
		ID:         createdOrderItem.ID,
		OrderID:    orderItemDTO.OrderID,
		MenuItemID: createdOrderItem.MenuItemID,
		Quantity:   createdOrderItem.Quantity,
		Subtotal:   createdOrderItem.Subtotal,
		// Set other fields accordingly
	}

	return createdOrderItemDTO, nil
}

func (c *orderItemUseCase) UpdateOrderItem(orderItemID uint, orderItemDTO *dtos.OrderItemDTO) (*dtos.OrderItemDTO, error) {
	orderItem := &model.OrderItem{
		OrderID:    orderItemDTO.OrderID,
		MenuItemID: orderItemDTO.MenuItemID,
		Quantity:   orderItemDTO.Quantity,
		Subtotal:   orderItemDTO.Subtotal,
		// Set other fields accordingly
	}

	updatedOrderItem, err := c.orderItemRepo.UpdateOrderItem(orderItem)
	if err != nil {
		return nil, err
	}

	updatedOrderItemDTO := &dtos.OrderItemDTO{
		ID:         updatedOrderItem.ID,
		OrderID:    orderItemDTO.OrderID,
		MenuItemID: updatedOrderItem.MenuItemID,
		Quantity:   updatedOrderItem.Quantity,
		Subtotal:   updatedOrderItem.Subtotal,
		// Set other fields accordingly
	}

	return updatedOrderItemDTO, nil
}

func (c *orderItemUseCase) RemoveOrderItem(orderItemID uint) error {
	err := c.orderItemRepo.DeleteOrderItem(orderItemID)
	if err != nil {
		return err
	}

	return nil
}

func (c *orderItemUseCase) GetOrderItemsByOrderID(orderID uint) ([]dtos.OrderItemDTO, error) {
	orderItems, err := c.orderItemRepo.GetOrderItemsByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	orderItemDTOs := make([]dtos.OrderItemDTO, len(orderItems))
	for i, orderItem := range orderItems {
		orderItemDTO := dtos.OrderItemDTO{
			ID:         orderItem.ID,
			OrderID:    orderItem.OrderID,
			MenuItemID: orderItem.MenuItemID,
			Quantity:   orderItem.Quantity,
			Subtotal:   orderItem.Subtotal,
			// Set other fields accordingly
		}
		orderItemDTOs[i] = orderItemDTO
	}

	return orderItemDTOs, nil
}
