package data

import (
	"altaproject2/domain"
	"altaproject2/features/product/data"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Userid  int
	Total   int
	Product []*data.Product `gorm:"many2many:product_orders;"`
}

type HistoryOrder struct {
	ProductName string
	Description string
	Price       int
	ProductPic  string
	Qty         int
	Subtotal    int
	Address     string
}

func (o *Order) ToModel() domain.Order {
	return domain.Order{
		Userid: o.Userid,
		Total:  o.Total,
	}
}

func (o *HistoryOrder) ToModelHistoryOrder() domain.HistoryOrder {
	return domain.HistoryOrder{
		ProductName: o.ProductName,
		Description: o.Description,
		Price:       o.Price,
		ProductPic:  o.ProductPic,
		Qty:         o.Qty,
		Subtotal:    o.Subtotal,
		Address:     o.Address,
	}
}

func (o *Order) ParseToArr(arr []HistoryOrder) []domain.HistoryOrder {
	var res []domain.HistoryOrder

	for _, val := range arr {
		res = append(res, val.ToModelHistoryOrder())
	}
	return res
}

func FromModel(data domain.Order) Order {
	var res Order
	res.Userid = data.Userid
	res.Total = data.Total
	return res
}
