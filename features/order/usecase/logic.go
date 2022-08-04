package usecase

import "altaproject2/domain"

type orderCase struct {
	orderData domain.OrderData
}

func New(od domain.OrderData) domain.OrderUseCase {
	return &orderCase{
		orderData: od,
	}
}

// DeleteOrder implements domain.OrderUseCase
func (*orderCase) DeleteOrder() {
	panic("unimplemented")
}

// GetOrder implements domain.OrderUseCase
func (*orderCase) GetOrder() {
	panic("unimplemented")
}

// PostOrder implements domain.OrderUseCase
func (*orderCase) PostOrder() {
	panic("unimplemented")
}
