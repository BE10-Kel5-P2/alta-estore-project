package data

import (
	"altaproject2/domain"

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
func (*cartData) PostData() {
	panic("unimplemented")
}

// UpdateData implements domain.CartData
func (*cartData) UpdateData() {
	panic("unimplemented")
}
