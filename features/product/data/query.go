package data

import (
	"altaproject2/domain"

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
func (*productData) PostItemData() {
	panic("unimplemented")
}

// UpdateItemData implements domain.ProductData
func (*productData) UpdateItemData() {
	panic("unimplemented")
}
