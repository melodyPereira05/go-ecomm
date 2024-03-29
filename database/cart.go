package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIsNotValid     = errors.New("user is not valid")
	ErrCantUpdateUser     = errors.New("cannot update user")
	ErrCantRemoveItemCart = errors.New("item cant be removed from cart")
	ErrCantGetItem        = errors.New("unable to find item")
	ErrCantBuyCartItem    = errors.New("cannot purchase the item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
