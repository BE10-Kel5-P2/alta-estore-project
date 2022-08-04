package domain

import "github.com/labstack/echo/v4"

type Order struct {
	ID      int
	Userid  int
	Total   int
	Product []Product
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

type OrderHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type OrderUseCase interface {
	PostOrder()
	GetOrder()
	DeleteOrder()
}

type OrderData interface {
	PostOrderData()
	GetOrderData()
	DeleteOrderData()
}
