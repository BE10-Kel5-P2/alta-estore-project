package usecase

import (
	"altaproject2/domain"

	"github.com/go-playground/validator/v10"
)

type productCase struct {
	productData domain.ProductData
	valid       *validator.Validate
}

func New(pd domain.ProductData, val *validator.Validate) domain.ProductUseCase {
	return &productCase{
		productData: pd,
		valid:       val,
	}
}

// DeleteItemAdmin implements domain.ProductUseCase
func (*productCase) DeleteItemAdmin() {
	panic("unimplemented")
}

// GetAllItems implements domain.ProductUseCase
func (pd *productCase) GetAllItems() ([]domain.Product, int) {
	get := pd.productData.GetAllItemData()

	if len(get) == 0 {
		return nil, 404
	}

	return get, 200
}

// GetItemUser implements domain.ProductUseCase
func (*productCase) GetItemUser() {
	panic("unimplemented")
}

// PostItemAdmin implements domain.ProductUseCase
func (*productCase) PostItemAdmin() {
	panic("unimplemented")
}

// UpdateItemAdmin implements domain.ProductUseCase
func (*productCase) UpdateItemAdmin() {
	panic("unimplemented")
}
