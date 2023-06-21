package usecase

import (
	"reglog/internal/dto/request"
	"reglog/internal/dto/response"
	"reglog/internal/model"
	"reglog/internal/repository"
)

type CartUseCase interface {
	AddItemToCart(request *request.AddCartItemRequest) (*response.AddCartItemResponse, error)
	UpdateCartItemQuantity(request *request.UpdateCartItemQuantityRequest) (*response.UpdateCartItemQuantityResponse, error)
	GetCartItems(cartID uint) (*response.GetCartItemsResponse, error)
	GetCartTotal(cartID uint) (*response.GetCartTotalResponse, error)
	RemoveItemFromCart(request *request.RemoveCartItemRequest) (*response.RemoveCartItemResponse, error)
	ClearCart(request *request.ClearCartRequest) (*response.ClearCartResponse, error)
}

type cartUseCase struct {
	cartRepo repository.CartRepository
	menuRepo repository.MenuRepository
}

func NewCartUseCase(cartRepo repository.CartRepository, menuRepo repository.MenuRepository) CartUseCase {
	return &cartUseCase{
		cartRepo: cartRepo,
		menuRepo: menuRepo,
	}
}

func (uc *cartUseCase) AddItemToCart(request *request.AddCartItemRequest) (*response.AddCartItemResponse, error) {
	cart, err := uc.cartRepo.GetCartByID(request.CartID)
	if err != nil {
		return nil, err
	}

	menuID := request.MenuID
	quantity := request.Quantity

	if uc.cartRepo.HasItemInCart(&cart, menuID) {
		// Item already exists in the cart, update the quantity
		cartItem, err := uc.cartRepo.GetCartItemByID(menuID)
		if err != nil {
			return nil, err
		}
		err = uc.cartRepo.UpdateCartItemQuantity(&cartItem, quantity)
		if err != nil {
			return nil, err
		}
	} else {
		// Item does not exist in the cart, create a new cart item
		menu, err := uc.menuRepo.GetMenuByID(menuID)
		if err != nil {
			return nil, err
		}
		subtotal := menu.Price * quantity
		cartItem := &model.CartItem{
			CartID: cart.ID,
			MenuID: menuID,
			Menu: model.Menu{
				Model:       menu.Model,
				Name:        menu.Name,
				Price:       menu.Price,
				Calories:    menu.Calories,
				City:        menu.City,
				Description: menu.Description,
				Ingredient:  menu.Ingredient,
				Images:      menu.Images,
				CategoryID:  menu.CategoryID,
				Category:    menu.Category,
			},
			Quantity: quantity,
			Subtotal: subtotal,
		}

		err = uc.cartRepo.AddItemToCart(&cart, cartItem)
		if err != nil {
			return nil, err
		}
	}

	response := &response.AddCartItemResponse{
		Message: "Item added to cart successfully",
	}
	return response, nil
}

func (uc *cartUseCase) UpdateCartItemQuantity(request *request.UpdateCartItemQuantityRequest) (*response.UpdateCartItemQuantityResponse, error) {
	cartItem, err := uc.cartRepo.GetCartItemByID(request.CartItemID)
	if err != nil {
		return nil, err
	}

	err = uc.cartRepo.UpdateCartItemQuantity(&cartItem, request.Quantity)
	if err != nil {
		return nil, err
	}

	response := &response.UpdateCartItemQuantityResponse{
		Message: "Cart item quantity updated successfully",
	}
	return response, nil
}

func (uc *cartUseCase) GetCartItems(cartID uint) (*response.GetCartItemsResponse, error) {
	cart, err := uc.cartRepo.GetCartByID(cartID)
	if err != nil {
		return nil, err
	}

	cartItems, err := uc.cartRepo.GetCartItems(&cart)
	if err != nil {
		return nil, err
	}

	var cartItemResponses []response.CartItemResponse
	for _, item := range cartItems {
		cartItemResponses = append(cartItemResponses, response.CartItemResponse{
			ItemID:   item.ID,
			MenuID:   item.MenuID,
			MenuName: item.Menu.Name, // Access the Name field from item.Menu
			Quantity: item.Quantity,
			Subtotal: item.Subtotal,
		})
	}

	response := &response.GetCartItemsResponse{
		CartItems: cartItemResponses,
	}

	return response, nil
}

func (uc *cartUseCase) GetCartTotal(cartID uint) (*response.GetCartTotalResponse, error) {
	cart, err := uc.cartRepo.GetCartByID(cartID)
	if err != nil {
		return nil, err
	}

	total, err := uc.cartRepo.GetCartTotal(&cart)
	if err != nil {
		return nil, err
	}

	response := &response.GetCartTotalResponse{
		Total: total,
	}

	return response, nil
}

func (uc *cartUseCase) RemoveItemFromCart(request *request.RemoveCartItemRequest) (*response.RemoveCartItemResponse, error) {
	cartItem, err := uc.cartRepo.GetCartItemByID(request.CartItemID)
	if err != nil {
		return nil, err
	}

	cart, err := uc.cartRepo.GetCartByID(cartItem.CartID)
	if err != nil {
		return nil, err
	}

	err = uc.cartRepo.RemoveItemFromCart(&cart, &cartItem)
	if err != nil {
		return nil, err
	}

	response := &response.RemoveCartItemResponse{
		Message: "Cart item removed successfully",
	}
	return response, nil
}

func (uc *cartUseCase) ClearCart(request *request.ClearCartRequest) (*response.ClearCartResponse, error) {
	cart, err := uc.cartRepo.GetCartByID(request.CartID)
	if err != nil {
		return nil, err
	}

	err = uc.cartRepo.ClearCart(&cart)
	if err != nil {
		return nil, err
	}

	response := &response.ClearCartResponse{
		Message: "Cart cleared successfully",
	}
	return response, nil
}
