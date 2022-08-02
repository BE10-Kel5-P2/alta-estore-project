package usecase

import (
	"altaproject2/domain"
	"altaproject2/features/product/data"

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
func (pd *productCase) PostItemAdmin(newProduct domain.Product) int {
	var product = data.FromModel(newProduct)
	validError := pd.valid.Struct(product)

	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}
	product.Qty = 0
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
