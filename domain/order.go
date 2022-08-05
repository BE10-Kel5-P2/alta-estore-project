package domain

import "github.com/labstack/echo/v4"

type Order struct {
	ID     int
	Userid int
	Total  int
	Data   []ProductOrders
}

type ProductOrders struct {
	ID        int
	Orderid   int
	Productid int
	Price     int
	Quantity  int
	Subtotal  int
	Status    int
	Payment   string
}

type OrderHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type OrderUseCase interface {
	PostOrder(neworder Order, data []ProductOrders) (int, string)
	Sum(neworder Order) int
	GetOrder(orderId int) (ProductOrders, error)
	DeleteOrder(userID, productID int) (bool, error)
}

type OrderData interface {
	PostOrderData(neworder Order) int
	PostProductOrderData(newpo []ProductOrders) []ProductOrders
	SumTotalPrice(neworder Order) int
	GetOrderData(oderId int) (ProductOrders, error)
	DeleteOrderData(userID, productID int) bool
}
