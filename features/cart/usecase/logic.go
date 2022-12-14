package usecase

import (
	"altaproject2/domain"
	"altaproject2/features/cart/data"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

type cartCase struct {
	cartData domain.CartData
	valid    *validator.Validate
}

func New(cd domain.CartData, val *validator.Validate) domain.CartUseCase {
	return &cartCase{
		cartData: cd,
		valid:    val,
	}
}

// Delete implements domain.CartUseCase
func (cd *cartCase) DeleteCart(userID int, productId int) (bool, error) {
	res := cd.cartData.DeleteData(userID, productId)

	if !res {
		return false, errors.New("failed to delete cart")
	}
	return true, nil
}

// GetCart implements domain.CartUseCase
func (cd *cartCase) GetCart(userid int) (domain.Cart, []domain.CartProduct, int) {
	getproduct := cd.cartData.GetDataProduct(userid)
	get := cd.cartData.GetData(userid)

	if get.ID == 0 {
		log.Println("data not found")
		return domain.Cart{}, nil, 404
	}

	return get, getproduct, 200
}

// PostCart implements domain.CartUseCase
func (cd *cartCase) PostCart(newcart domain.Cart) int {
	var cartdata = data.FromModel(newcart)

	validError := cd.valid.Struct(cartdata)

	if validError != nil {
		log.Println("Validation errror : ", validError)
		return 400
	}

	duplicate := cd.cartData.CheckDuplicate(cartdata.ToModel())

	if duplicate {
		log.Println("Duplicate Data")
		return 400
	}

	insert := cd.cartData.PostData(cartdata.ToModel())

	if insert.ID == 0 {
		log.Println("Empty data")
		return 404
	}

	return 200
}

// UpdateCart implements domain.CartUseCase
func (cc *cartCase) UpdateData(newUpdate domain.Cart, productId int) int {
	var cart = data.FromModel(newUpdate)

	if productId == 0 {
		log.Println("Data not found")
		return 404
	}

	cart.Productid = int(productId)

	update := cc.cartData.UpdateData(cart.ToModel(), productId)

	if update.ID == 0 {
		log.Println("Data not found")
		return 404
	}

	return 200
}
