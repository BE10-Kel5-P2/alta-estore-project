package data

import (
	"altaproject2/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Userid int
	Total  int
	Data   []ProductOrders `json:"data" gorm:"foreignKey:Orderid"`
}

type ProductOrders struct {
	gorm.Model
	Orderid   int
	Productid int `json:"productid" form:"productid"`
	Price     int `json:"price" form:"price"`
	Quantity  int `json:"quantity" form:"quantity"`
	Subtotal  int `json:"subtotal" form:"subtotal"`
	Status    int `json:"status" form:"status"`
	Payment   string
}

func (o *Order) ToModel() domain.Order {
	return domain.Order{
		Userid: o.Userid,
		Total:  o.Total,
	}
}

func (o *ProductOrders) ToModelProductOrders() domain.ProductOrders {
	return domain.ProductOrders{
		Orderid:   o.Orderid,
		Productid: o.Orderid,
		Price:     o.Price,
		Quantity:  o.Quantity,
		Subtotal:  o.Subtotal,
		Status:    o.Status,
	}
}

func ParseToArr(arr []Order) []domain.Order {
	var res []domain.Order

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.Order) Order {
	var res Order
	res.Userid = data.Userid
	res.Total = data.Total
	return res
}
