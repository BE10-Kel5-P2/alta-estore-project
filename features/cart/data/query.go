package data

import (
	"altaproject2/domain"
	"log"

	"gorm.io/gorm"
)

func New(db *gorm.DB) domain.CartData {
	return &cartData{
		db: db,
	}
}

type cartData struct {
	db *gorm.DB
}

// DeleteData implements domain.CartData
func (*cartData) DeleteData() {
	panic("unimplemented")
}

// GetData implements domain.CartData
func (*cartData) GetData() {
	panic("unimplemented")
}

// PostData implements domain.CartData
func (cd *cartData) PostData(newcart domain.Cart) domain.Cart {
	var cartdata = FromModel(newcart)
	err := cd.db.Create(&cartdata).Error

	if err != nil {
		log.Println("Cant create user object", err.Error())
		return domain.Cart{}
	}

	return cartdata.ToModel()
}

// UpdateData implements domain.CartData
func (cd *cartData) UpdateData(newUpdate domain.Cart, productID int) domain.Cart {
	var cart = FromModel(newUpdate)
	err := cd.db.Model(&Cart{}).Where("Productid = ?", productID).Updates(cart)

	if err.Error != nil {
		log.Println("Cant update cart object", err.Error.Error())
		return domain.Cart{}
	}

	if err.RowsAffected == 0 {
		log.Println("Data Not Found")
		return domain.Cart{}
	}

	return cart.ToModel()
}
