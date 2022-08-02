package data

import (
	"altaproject2/domain"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Subtotal  int
	Productid int
}

func (u *Cart) ToModel() domain.Cart {
	return domain.Cart{
		ID:        int(u.ID),
		Subtotal:  u.Subtotal,
		Productid: u.Productid,
	}
}

func ParseToArr(arr []Cart) []domain.Cart {
	var res []domain.Cart

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Cart) Cart {
	var res Cart
	res.ID = uint(data.ID)
	res.Subtotal = data.Subtotal
	res.Productid = data.Subtotal
	return res
}
