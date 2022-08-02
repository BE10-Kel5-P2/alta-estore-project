package domain

import "github.com/labstack/echo/v4"

type Cart struct {
	ID        int
	Subtotal  int
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
	UpdateCart()
	GetCart()
	DeleteCart()
}

type CartData interface {
	PostData()
	UpdateData()
	GetData()
	DeleteData()
}
