package repository

import (
	"reglog/internal/model"

	"github.com/jinzhu/gorm"
)

type CartRepository interface {
	GetCartByID(id uint) (model.Cart, error)
	AddItemToCart(cart *model.Cart, cartItem *model.CartItem) error
	RemoveItemFromCart(cart *model.Cart, cartItem *model.CartItem) error
	UpdateCartItemQuantity(cartItem *model.CartItem, quantity float64) error
	GetCartItems(cart *model.Cart) ([]model.CartItem, error)
	ClearCart(cart *model.Cart) error
	GetCartTotal(cart *model.Cart) (float64, error)
	SetCartItemQuantity(cartItem *model.CartItem, quantity float64) error
	HasItemInCart(cart *model.Cart, menuID uint) bool
	GetCartItemByID(id uint) (model.CartItem, error)
	UpdateCartItems(cartItems []model.CartItem) error
	SaveCart(cart *model.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetCartByID(id uint) (model.Cart, error) {
	var cart model.Cart
	err := r.db.Preload("Items").Where("id = ?", id).First(&cart).Error
	if err != nil {
		return model.Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) AddItemToCart(cart *model.Cart, cartItem *model.CartItem) error {
	return r.db.Model(cart).Association("Items").Append(cartItem).Error
}

func (r *cartRepository) RemoveItemFromCart(cart *model.Cart, cartItem *model.CartItem) error {
	return r.db.Model(cart).Association("Items").Delete(cartItem).Error
}

func (r *cartRepository) UpdateCartItemQuantity(cartItem *model.CartItem, quantity float64) error {
	cartItem.Quantity = quantity
	return r.db.Save(cartItem).Error
}

func (r *cartRepository) GetCartItems(cart *model.Cart) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	err := r.db.Model(cart).Association("Items").Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (r *cartRepository) ClearCart(cart *model.Cart) error {
	return r.db.Model(cart).Association("Items").Clear().Error
}

func (r *cartRepository) GetCartTotal(cart *model.Cart) (float64, error) {
	var total float64
	err := r.db.Model(cart).Select("SUM(subtotal)").Scan(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *cartRepository) SetCartItemQuantity(cartItem *model.CartItem, quantity float64) error {
	cartItem.Quantity = quantity
	cartItem.Subtotal = cartItem.Menu.Price * quantity
	return r.db.Save(cartItem).Error
}

func (r *cartRepository) HasItemInCart(cart *model.Cart, menuID uint) bool {
	var count int
	r.db.Where("cart_id = ? AND menu_id = ?", cart.ID, menuID).Model(&model.CartItem{}).Count(&count)
	return count > 0
}

func (r *cartRepository) GetCartItemByID(id uint) (model.CartItem, error) {
	var cartItem model.CartItem
	err := r.db.Where("id = ?", id).First(&cartItem).Error
	if err != nil {
		return model.CartItem{}, err
	}
	return cartItem, nil
}

func (r *cartRepository) UpdateCartItems(cartItems []model.CartItem) error {
	return r.db.Save(cartItems).Error
}

func (r *cartRepository) SaveCart(cart *model.Cart) error {
	return r.db.Save(cart).Error
}
