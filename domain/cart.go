package domain

import "github.com/labstack/echo/v4"

type Cart struct {
	ID        int
	Subtotal  int
	Userid    int
	Productid int
	Quantity  int
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
	UpdateData(newUpdate Cart, productId int) int
	PostCart(newcart Cart) int
	GetCart()
	DeleteCart(productID int) (bool, error)
}

type CartData interface {
	UpdateData(newUpdate Cart, productId int) Cart
	PostData(newcart Cart) Cart
	GetData()
	DeleteData(productID int) bool
}
