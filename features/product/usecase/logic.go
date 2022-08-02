package usecase

import (
	"altaproject2/domain"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
func (pd *productCase) GetItemUser(id int) (domain.Product, error) {
	data, err := pd.productData.GetItemData(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.Product{}, errors.New("data not found")
		} else {
			return domain.Product{}, errors.New("server error")
		}
	}

	return data, nil
}

// PostItemAdmin implements domain.ProductUseCase
func (*productCase) PostItemAdmin() {
	panic("unimplemented")
}

// UpdateItemAdmin implements domain.ProductUseCase
func (*productCase) UpdateItemAdmin() {
	panic("unimplemented")
}
