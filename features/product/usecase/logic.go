package usecase

import (
	"altaproject2/domain"
	"altaproject2/features/product/data"
	"log"

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
func (*productCase) GetAllItems() {
	panic("unimplemented")
}

// GetItemUser implements domain.ProductUseCase
func (*productCase) GetItemUser() {
	panic("unimplemented")
}

// PostCartUser implements domain.ProductUseCase
func (*productCase) PostCartUser() {
	panic("unimplemented")
}

// PostItemAdmin implements domain.ProductUseCase
func (pd *productCase) PostItemAdmin(newProduct domain.Product) int {
	var product = data.FromModel(newProduct)
	validError := pd.valid.Struct(product)

	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	insert := pd.productData.PostItemData(product.ToModel())

	if insert.ID == 0 {
		log.Println("Empty data")
		return 404
	}

	return 200
}

// UpdateItemAdmin implements domain.ProductUseCase
func (pd *productCase) UpdateItemAdmin(newProduct domain.Product, productID int) int {
	var product = data.FromModel(newProduct)

	if productID == 0 {
		log.Println("Data not found")
		return 404
	}

	product.ID = uint(productID)

	pd.productData.UpdateItemData(product.ToModel(), productID)

	return 200
}
