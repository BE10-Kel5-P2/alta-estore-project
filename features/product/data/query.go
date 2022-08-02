package data

import (
	"altaproject2/domain"
	"log"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductData {
	return &productData{
		db: db,
	}
}

// DeleteItemData implements domain.ProductData
func (*productData) DeleteItemData() {
	panic("unimplemented")
}

// GetAllItemData implements domain.ProductData
func (*productData) GetAllItemData() {
	panic("unimplemented")
}

// GetItemData implements domain.ProductData
func (*productData) GetItemData() {
	panic("unimplemented")
}

// PostCartData implements domain.ProductData
func (*productData) PostCartData() {
	panic("unimplemented")
}

// PostItemData implements domain.ProductData
func (pd *productData) PostItemData(newProduct domain.Product) domain.Product {
	var product = FromModel(newProduct)
	err := pd.db.Create(&product).Error

	if product.ID == 0 {
		log.Println("Invalid ID")
		return domain.Product{}
	}

	if err != nil {
		log.Println("Cant create product object", err.Error())
		return domain.Product{}
	}

	return product.ToModel()
}

// UpdateItemData implements domain.ProductData
func (pd *productData) UpdateItemData() {
	panic("unimplemented")
}
