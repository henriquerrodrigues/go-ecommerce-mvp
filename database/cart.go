package database

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantRemoveItemCart = errors.New("cannot add this product to the cart")
	ErrCantGetItem        = errors.New("cannot remove this item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func addProductToCart() gin.HandlerFunc {

}

func RemoveCartItem() gin.HandlerFunc {

}

func BuyItemFromCart() gin.HandlerFunc {

}

func InstantBuyer() gin.HandlerFunc {

}
