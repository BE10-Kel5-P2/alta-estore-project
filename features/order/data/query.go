package data

import (
	"altaproject2/domain"

	"gorm.io/gorm"
)

type orderData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OrderData {
	return &orderData{
		db: db,
	}
}

// DeleteOrderData implements domain.OrderData
func (*orderData) DeleteOrderData() {
	panic("unimplemented")
}

// GetOrderData implements domain.OrderData
func (*orderData) GetOrderData() {
	panic("unimplemented")
}

// PostOrderData implements domain.OrderData
func (*orderData) PostOrderData() {
	panic("unimplemented")
}
