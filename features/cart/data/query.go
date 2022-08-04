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
func (cd *cartData) DeleteData(userID int, productID int) bool {
	err := cd.db.Where("userID = ? AND productID = ?", userID, productID).Delete(&Cart{})

	if err.Error != nil {
		log.Println("cannot delete data", err.Error.Error())
		return false
	}

	if err.RowsAffected < 1 {
		log.Println("No content deleted", err.Error.Error())
		return false
	}

	return true
}

// CheckDuplicate implements domain.CartData
func (cd *cartData) CheckDuplicate(newcart domain.Cart) bool {
	var cart Cart
	err := cd.db.Find(&cart, "productid = ? AND deleted_at IS NULL", newcart.Productid)

	if err.RowsAffected == 1 {
		log.Println("Duplicated data", err.Error)
		return true
	}

	return false
}

// GetDataProduct implements domain.CartData
func (cd *cartData) GetDataProduct(userid int) []domain.CartProduct {
	var cproduct []CartProduct

	err := cd.db.Model(&Cart{}).Select("*").Joins("join products on products.id = carts.productid").Where("userid = ?", userid).Scan(&cproduct)

	if err.Error != nil {
		log.Println("Cannot retrieve object", err.Error)
		return nil
	}

	return ParseToArrProduct(cproduct)
}

// GetData implements domain.CartData
func (cd *cartData) GetData(userid int) domain.Cart {
	var cart domain.Cart

	err := cd.db.Find(&cart, "userid = ?", userid)

	if err.Error != nil {
		log.Println("cant get all data from db", err.Error)
		return domain.Cart{}
	}

	if err.RowsAffected == 0 {
		log.Println("data not found")
		return domain.Cart{}
	}

	return cart
}

// PostData implements domain.CartData
func (cd *cartData) PostData(newcart domain.Cart) domain.Cart {
	var cartdata = FromModel(newcart)
	var subtotal int

	cd.db.Model(&Cart{}).Select("? * products.price", cartdata.Quantity).Joins("join products on products.id = ?", cartdata.Productid).Scan(&subtotal)
	cartdata.Subtotal = subtotal
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
