package controllers

import "errors"

var (
	ErrCantFindProduct          = errors.New("Cant find the Product")
	ErrCantDecodeProducts       = errors.New("Cant find the Product")
	ErrUserIdIsNotValid         = errors.New("This User is not valid")
	ErrCantUpdateUser           = errors.New("Cannot add this product to the cart")
	ErrCantRemoteRemoveItemCart = errors.New("Cannot remove this item from the cart")
	ErrCantGetItem              = errors.New("Cant get the item")
	ErrCantBuyCartItem          = errors.New("Cant buy the item from the cart")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
