package controllers

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("Cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("Cannot remove item from the cart")
	ErrCanGetItem         = errors.New("Was unable to get the item from the cart")
	ErrCanBuyCartItem     = errors.New("Cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
