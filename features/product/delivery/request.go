package delivery

import "altaproject2/domain"

type ProductFormat struct {
	ProductName string `json:"productname" form:"productname" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	ProductPic  string `json:"productpic" form:"productpic" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Qty         int    `json:"qty" form:"qty" validate:"required"`
}

func (i *ProductFormat) ToModel() domain.Product {
	return domain.Product{
		ProductName: i.ProductName,
		Description: i.Description,
		Price:       i.Price,
		ProductPic:  i.ProductPic,
		Stock:       i.Stock,
		Qty:         i.Qty,
	}
}
