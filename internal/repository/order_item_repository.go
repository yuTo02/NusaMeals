package repository

import (
	"github.com/jinzhu/gorm"
	"reglog/internal/model"
)

type OrderItemRepository interface {
	CreateOrderItem(orderItem *model.OrderItem) (*model.OrderItem, error)
	UpdateOrderItem(orderItem *model.OrderItem) (*model.OrderItem, error)
	DeleteOrderItem(orderItemID uint) error
	GetOrderItemByID(orderItemID uint) (*model.OrderItem, error)
	GetAllOrderItems() ([]model.OrderItem, error)
	GetOrderItemsByOrderID(orderID uint) ([]model.OrderItem, error)
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{
		db: db,
	}
}

func (r *orderItemRepository) CreateOrderItem(orderItem *model.OrderItem) (*model.OrderItem, error) {
	err := r.db.Create(orderItem).Error
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (r *orderItemRepository) UpdateOrderItem(orderItem *model.OrderItem) (*model.OrderItem, error) {
	err := r.db.Save(orderItem).Error
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (r *orderItemRepository) DeleteOrderItem(orderItemID uint) error {
	return r.db.Delete(&model.OrderItem{}, orderItemID).Error
}

func (r *orderItemRepository) GetOrderItemByID(orderItemID uint) (*model.OrderItem, error) {
	var orderItem model.OrderItem
	err := r.db.First(&orderItem, orderItemID).Error
	if err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (r *orderItemRepository) GetAllOrderItems() ([]model.OrderItem, error) {
	var orderItems []model.OrderItem
	err := r.db.Find(&orderItems).Error
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *orderItemRepository) GetOrderItemsByOrderID(orderID uint) ([]model.OrderItem, error) {
	var orderItems []model.OrderItem
	err := r.db.Where("order_id = ?", orderID).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}
