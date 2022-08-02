package usecase

import "altaproject2/domain"

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
func (*cartCase) UpdateCart() {
	panic("unimplemented")
}
