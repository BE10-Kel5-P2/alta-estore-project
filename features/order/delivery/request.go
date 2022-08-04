package delivery

import (
	"altaproject2/domain"
)

type OrderFormat struct {
	Userid  int
	Total   int
	Payment string
	Data    []ProductOrdersFormat `json:"data"`
}

type ProductOrdersFormat struct {
	Productid int `json:"productid" form:"productid"`
	Price     int `json:"price" form:"price"`
	Quantity  int `json:"quantity" form:"quantity"`
	Subtotal  int `json:"subtotal" form:"subtotal"`
	Status    int `json:"status" form:"status"`
}

func (i *OrderFormat) ToModel() domain.Order {
	return domain.Order{
		Userid: i.Userid,
		Total:  i.Total,
	}
}

func (i *ProductOrdersFormat) ToModelProductOrders() domain.ProductOrders {
	return domain.ProductOrders{
		Productid: i.Productid,
		Price:     i.Price,
		Quantity:  i.Quantity,
		Subtotal:  i.Subtotal,
		Status:    i.Status,
	}
}
