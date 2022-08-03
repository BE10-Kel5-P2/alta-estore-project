package domain

import "github.com/labstack/echo/v4"

type Cart struct {
	ID        int
	Subtotal  int
	Qty       int
	Productid int
}

type CartHandler interface {
	Post() echo.HandlerFunc
	Update() echo.HandlerFunc
	Get() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CartUseCase interface {
	PostCart()
	UpdateData(newUpdate Cart, productId int) int
	GetCart()
	DeleteCart()
}

type CartData interface {
	PostData()
	UpdateData(newUpdate Cart, productId int) Cart
	GetData()
	DeleteData()
}
