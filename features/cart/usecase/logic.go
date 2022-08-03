package usecase

import (
	"altaproject2/domain"
	"altaproject2/features/cart/data"
	"log"
)

type cartCase struct {
	cartData domain.CartData
}

func New(pd domain.CartData) domain.CartUseCase {
	return &cartCase{
		cartData: pd,
	}
}

// Delete implements domain.CartUseCase
func (*cartCase) DeleteCart() {
	panic("unimplemented")
}

// GetCart implements domain.CartUseCase
func (*cartCase) GetCart() {
	panic("unimplemented")
}

// PostCart implements domain.CartUseCase
func (*cartCase) PostCart() {
	panic("unimplemented")
}

// UpdateCart implements domain.CartUseCase
func (cc *cartCase) UpdateData(newUpdate domain.Cart, productId int) int {
	var cart = data.FromModel(newUpdate)

	if productId == 0 {
		log.Println("Data not found")
		return 404
	}

	cart.Productid = int(productId)

	cc.cartData.UpdateData(cart.ToModel(), productId)

	return 200
}
