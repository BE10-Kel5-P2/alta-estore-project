package data

import (
	"altaproject2/domain"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Userid    int
	Productid int `json:"productid" form:"productid" validate:"required"`
	Quantity  int `json:"quantity" form:"quantity" validate:"required"`
	Subtotal  int
}

func (u *Cart) ToModel() domain.Cart {
	return domain.Cart{
		ID:        int(u.ID),
		Userid:    u.Userid,
		Productid: u.Productid,
		Quantity:  u.Quantity,
		Subtotal:  u.Subtotal,
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
	res.Userid = data.Userid
	res.Productid = data.Productid
	res.Quantity = data.Quantity
	res.Subtotal = data.Subtotal
	return res
}
