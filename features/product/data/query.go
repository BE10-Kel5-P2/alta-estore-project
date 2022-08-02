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
func (pd *productData) GetAllItemData() []domain.Product {
	var products []Product

	err := pd.db.Find(&products)

	if err.Error != nil {
		log.Println("cant get all data from db", err.Error)
		return nil
	}

	return ParseToArr(products)
}

// GetItemData implements domain.ProductData
func (*productData) GetItemData() {
	panic("unimplemented")
}

// PostItemData implements domain.ProductData
func (*productData) PostItemData() {
	panic("unimplemented")
}

// UpdateItemData implements domain.ProductData
func (*productData) UpdateItemData() {
	panic("unimplemented")
}
