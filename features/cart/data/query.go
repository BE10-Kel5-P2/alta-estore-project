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
func (cd *cartData) DeleteData(productID int) bool {
	err := cd.db.Where("ID = ?", productID).Delete(&Cart{})

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
