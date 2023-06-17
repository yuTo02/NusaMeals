package usecase

import (
	"reglog/internal/dtos"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type OrderUseCase interface {
	CreateOrder(orderDTO *dtos.OrderDTO) (*dtos.OrderDTO, error)
	GetOrderByID(orderID uint) (*dtos.OrderDTO, error)
}

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository) OrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
	}
}

func (c *orderUseCase) CreateOrder(orderDTO *dtos.OrderDTO) (*dtos.OrderDTO, error) {
	orderItems := make([]model.OrderItem, len(orderDTO.OrderItems))
	for i, item := range orderDTO.OrderItems {
		orderItems[i] = model.OrderItem{
			// Assign nilai-nilai yang sesuai dari orderDTO ke OrderItem
			MenuItemID: item.MenuItemID,
			Quantity:   item.Quantity,
			Subtotal:   item.Subtotal,
			// ...
		}
	}
	order := &model.Order{
		UserID:      orderDTO.UserID,
		OrderItems:  orderItems,
		TotalAmount: orderDTO.TotalAmount,
		// Set other fields accordingly
	}

	createdOrder, err := c.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	createdOrderDTO := &dtos.OrderDTO{
		ID:     createdOrder.ID,
		UserID: createdOrder.UserID,
		// Set other fields accordingly
	}

	return createdOrderDTO, nil
}

func (c *orderUseCase) GetOrderByID(orderID uint) (*dtos.OrderDTO, error) {
	order, err := c.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	orderDTO := &dtos.OrderDTO{
		ID:     order.ID,
		UserID: order.UserID,
		// Set other fields accordingly
	}

	return orderDTO, nil
}
