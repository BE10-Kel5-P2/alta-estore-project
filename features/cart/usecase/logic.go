package usecase

import (
	"altaproject2/domain"
	"errors"
)

type cartCase struct {
	cartData domain.CartData
}

func New(cd domain.CartData) domain.CartUseCase {
	return &cartCase{
		cartData: cd,
	}
}

// Delete implements domain.CartUseCase
func (cd *cartCase) DeleteCart(productId int) (bool, error) {
	res := cd.cartData.DeleteData(productId)

	if !res {
		return false, errors.New("failed to delete cart")
	}
	return true, nil
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
func (*cartCase) UpdateCart() {
	panic("unimplemented")
}
