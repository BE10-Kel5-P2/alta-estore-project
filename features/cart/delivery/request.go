package delivery

import "altaproject2/domain"

type CartFormat struct {
	Subtotal  int `json:"subtotal" form:"subtotal" validate:"required"`
	Qty       int `json:"qty" form:"qty" validate:"required"`
	Productid int `json:"productid" form:"productid" validate:"required"`
}

func (i *CartFormat) ToModel() domain.Cart {
	return domain.Cart{
		Subtotal:  i.Subtotal,
		Qty:       i.Qty,
		Productid: i.Productid,
	}
}
