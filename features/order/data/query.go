package data

import (
	"altaproject2/domain"
	"log"

	"gorm.io/gorm"
)

type orderData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OrderData {
	return &orderData{
		db: db,
	}
}

// SumTotalPrice implements domain.OrderData
func (od *orderData) SumTotalPrice(neworder domain.Order) int {
	total := 0

	res := od.db.Model(&domain.Cart{}).Select("SUM(carts.subtotal)").Joins("join users on users.id = carts.userid").
		Where("users.id = ?", neworder.Userid).Scan(&total)

	if res.Error != nil {
		log.Println("Cant get total", res.Error)
		return 0
	}

	return total
}

// PostProductOrderData implements domain.OrderData
func (od *orderData) PostProductOrderData(newpo []domain.ProductOrders) []domain.ProductOrders {

	err := od.db.Create(&newpo)

	if err.Error != nil {
		log.Println("Cant send data")
		return nil
	}

	od.db.Exec("update product_orders join products on product_orders.productid = products.id set products.stock = products.stock - product_orders.quantity where product_orders.productid = 1")
	od.db.Raw("Truncate table carts")
	return newpo
}

// DeleteOrderData implements domain.OrderData
func (od *orderData) DeleteOrderData(userID int, productID int) bool {
	err := od.db.Where("userID = ? AND ID = ?", userID, productID).Delete(&Order{})

	if err.Error != nil {
		log.Println("cannot delete data", err.Error.Error())
		return false
	}

	if err.RowsAffected < 1 {
		log.Println("No content deleted", err.Error.Error())
		return false
	}
	od.db.Exec("update product_orders join products on product_orders.productid = products.id set products.stock = products.stock + product_orders.quantity where product_orders.productid = 1")
	return true
}

// GetOrderData implements domain.OrderData
func (od *orderData) GetOrderData(orderId int) (domain.ProductOrders, error) {
	var tmp ProductOrders
	err := od.db.Where("ID = ?", orderId).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.ProductOrders{}, err
	}

	return tmp.ToModelProductOrders(), nil
}

// PostOrderData implements domain.OrderData
func (od *orderData) PostOrderData(neworder domain.Order) int {
	var order = FromModel(neworder)

	err := od.db.Create(&order).Error

	if err != nil {
		log.Println("Cant send data")
		return 0
	}

	return int(order.ID)
}
