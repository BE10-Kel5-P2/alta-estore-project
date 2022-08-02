package data

import (
	"altaproject2/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `json:"productname" form:"productname" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	ProductPic  string `json:"productpic" form:"productpic" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Qty         int    `json:"qty" form:"qty"`
}

func (u *Product) ToModel() domain.Product {
	return domain.Product{
		ID:          int(u.ID),
		ProductName: u.ProductName,
		Description: u.Description,
		Price:       u.Price,
		ProductPic:  u.ProductPic,
		Stock:       u.Stock,
		Qty:         u.Qty,
	}
}

func ParseToArr(arr []Product) []domain.Product {
	var res []domain.Product

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.Product) Product {
	var res Product
	res.ID = uint(data.ID)
	res.ProductName = data.ProductName
	res.Description = data.Description
	res.Price = data.Price
	res.ProductPic = data.ProductPic
	res.Stock = data.Stock
	res.Qty = data.Qty
	return res
}
