package domain

import "github.com/labstack/echo/v4"

type Cart struct {
	ID        int
	Subtotal  int
	Productid int
}

type ProductCart struct {
	ProductName string
	Price       int
	ProductPic  string
	Stock       int
	Qty         int
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
	DeleteCart(productID int) (bool, error)
}

type CartData interface {
	PostData()
	UpdateData()
	GetData()
	DeleteData(productID int) bool
}
