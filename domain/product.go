package domain

import "github.com/labstack/echo/v4"

type Product struct {
	ID          int
	ProductName string
	Description string
	Price       int
	ProductPic  string
	Stock       int
	Qty         int
	// Cart []Cart
}

type ProductHandler interface {
	GetItems() echo.HandlerFunc
	PostItem() echo.HandlerFunc
	UpdateItem() echo.HandlerFunc
	DeleteItem() echo.HandlerFunc
	GetItem() echo.HandlerFunc
}

type ProductUseCase interface {
	GetAllItems() ([]Product, int)
	PostItemAdmin()
	UpdateItemAdmin()
	DeleteItemAdmin()
	GetItemUser()
}

type ProductData interface {
	GetAllItemData() []Product
	PostItemData()
	UpdateItemData()
	DeleteItemData()
	GetItemData()
}
