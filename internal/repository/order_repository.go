package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	CreateOrder(order *model.Order) (*model.Order, error)
	UpdateOrder(order *model.Order) (*model.Order, error)
	DeleteOrder(orderID uint) error
	GetOrderByID(orderID uint) (*model.Order, error)
	GetAllOrders() ([]model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) CreateOrder(order *model.Order) (*model.Order, error) {
	err := r.db.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) UpdateOrder(order *model.Order) (*model.Order, error) {
	err := r.db.Save(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) DeleteOrder(orderID uint) error {
	order := &model.Order{}
	order.ID = orderID
	return r.db.Delete(order).Error
}

func (r *orderRepository) GetOrderByID(orderID uint) (*model.Order, error) {
	var order model.Order
	err := r.db.First(&order, orderID).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
