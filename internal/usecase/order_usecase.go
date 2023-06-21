package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type OrderUseCase interface {
	CreateOrder(orderDTO *request.Order) (*response.Order, error)
	GetOrderByID(orderID uint) (*response.Order, error)
	UpdateOrder(orderDTO *request.OrderUpdate) (*response.Order, error)
	DeleteOrder(orderID uint) error
	GetAllOrders() ([]response.Order, error)
}

type orderUseCase struct {
	orderRepo repository.OrderRepository
	menuRepo  repository.MenuRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository, menuRepo repository.MenuRepository) OrderUseCase {
	return &orderUseCase{
		orderRepo: orderRepo,
		menuRepo:  menuRepo,
	}
}

func (c *orderUseCase) CreateOrder(orderDTO *request.Order) (*response.Order, error) {
	menu, err := c.menuRepo.GetMenuByID(orderDTO.MenuID)
	if err != nil {
		return nil, err
	}

	quantity := int(orderDTO.Quantity)
	totalPrice := menu.Price * quantity

	order := &model.Order{
		UserID:      orderDTO.UserID,
		MenuID:      orderDTO.MenuID,
		Quantity:    orderDTO.Quantity,
		TypeOrder:   orderDTO.TypeOrder,
		OrderStatus: "on progress",
		TotalPrice:  totalPrice,
		// Set other fields accordingly
	}

	createdOrder, err := c.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	createdOrderResponse := &response.Order{
		ID:          createdOrder.ID,
		UserID:      createdOrder.UserID,
		MenuID:      createdOrder.MenuID,
		Quantity:    createdOrder.Quantity,
		TotalPrice:  createdOrder.TotalPrice,
		TypeOrder:   createdOrder.TypeOrder,
		OrderStatus: createdOrder.OrderStatus,
		CreatedAt:   createdOrder.CreatedAt,
		UpdatedAt:   createdOrder.UpdatedAt,
	}

	return createdOrderResponse, nil
}

func (c *orderUseCase) GetOrderByID(orderID uint) (*response.Order, error) {
	order, err := c.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	orderResponse := &response.Order{
		ID:          order.ID,
		UserID:      order.UserID,
		MenuID:      order.MenuID,
		Quantity:    order.Quantity,
		TypeOrder:   order.TypeOrder,
		TotalPrice:  order.TotalPrice,
		OrderStatus: order.OrderStatus,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}

	return orderResponse, nil
}

func (c *orderUseCase) UpdateOrder(orderDTO *request.OrderUpdate) (*response.Order, error) {
	order, err := c.orderRepo.GetOrderByID(orderDTO.ID)
	if err != nil {
		return nil, err
	}

	// Update fields accordingly
	order.OrderStatus = orderDTO.OrderStatus
	// Set other fields accordingly

	updatedOrder, err := c.orderRepo.UpdateOrder(order)
	if err != nil {
		return nil, err
	}

	updatedOrderResponse := &response.Order{
		ID:          updatedOrder.ID,
		UserID:      updatedOrder.UserID,
		MenuID:      updatedOrder.MenuID,
		Quantity:    updatedOrder.Quantity,
		TypeOrder:   updatedOrder.TypeOrder,
		TotalPrice:  updatedOrder.TotalPrice,
		OrderStatus: updatedOrder.OrderStatus,
		CreatedAt:   updatedOrder.CreatedAt,
		UpdatedAt:   updatedOrder.UpdatedAt,
	}

	return updatedOrderResponse, nil
}

func (c *orderUseCase) DeleteOrder(orderID uint) error {
	err := c.orderRepo.DeleteOrder(orderID)
	if err != nil {
		return err
	}

	return nil
}

func (c *orderUseCase) GetAllOrders() ([]response.Order, error) {
	orders, err := c.orderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	orderResponses := make([]response.Order, len(orders))
	for i, order := range orders {
		orderResponses[i] = response.Order{
			ID:          order.ID,
			UserID:      order.UserID,
			MenuID:      order.MenuID,
			Quantity:    order.Quantity,
			TypeOrder:   order.TypeOrder,
			TotalPrice:  order.TotalPrice,
			OrderStatus: order.OrderStatus,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		}
	}

	return orderResponses, nil
}
