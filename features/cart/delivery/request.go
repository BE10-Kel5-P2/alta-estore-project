package delivery

import "altaproject2/domain"

type CartFormat struct {
	Userid    int
	Productid int `json:"productid" form:"productid"`
	Quantity  int `json:"quantity" form:"quantity"`
	Subtotal  int
}

func (i *CartFormat) ToModel() domain.Cart {
	return domain.Cart{
		Userid:    i.Userid,
		Productid: i.Productid,
		Quantity:  i.Quantity,
		Subtotal:  i.Subtotal,
	}
}
