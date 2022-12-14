package usecase

import (
	"altaproject2/domain"
	"altaproject2/features/order/data"
	midconn "altaproject2/utils/midtrans"
	"errors"
	"log"

	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type orderCase struct {
	orderData domain.OrderData
	midconn   snap.Client
}

func New(od domain.OrderData, mc snap.Client) domain.OrderUseCase {
	return &orderCase{
		orderData: od,
		midconn:   mc,
	}
}

// Sum implements domain.OrderUseCase
func (od *orderCase) Sum(neworder domain.Order) int {
	total := od.orderData.SumTotalPrice(neworder)
	return total
}

// DeleteOrder implements domain.OrderUseCase
func (od *orderCase) DeleteOrder(userID int, productId int) (bool, error) {
	res := od.orderData.DeleteOrderData(userID, productId)

	if !res {
		return false, errors.New("failed delete data")
	}
	return true, nil

}

// GetOrder implements domain.OrderUseCase
func (od *orderCase) GetOrder(orderId int) (domain.ProductOrders, error) {
	data, err := od.orderData.GetOrderData(orderId)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.ProductOrders{}, errors.New("data not found")
		} else {
			return domain.ProductOrders{}, errors.New("server error")
		}
	}

	return data, nil
}

// PostOrder implements domain.OrderUseCase
func (od *orderCase) PostOrder(neworder domain.Order, datapo []domain.ProductOrders) (int, string) {
	var order = data.FromModel(neworder)
	id := od.orderData.PostOrderData(order.ToModel())

	if id == 0 {
		log.Println("data not found")
		return 404, ""
	}

	res := midconn.CreateConnection(od.midconn, order.Total, id)

	for i := 0; i < len(datapo); i++ {
		datapo[i].Orderid = id
		datapo[i].Status = 0
		datapo[i].Payment = res.RedirectURL
	}

	prodorder := od.orderData.PostProductOrderData(datapo)

	if prodorder == nil {
		log.Println("data not found")
		return 404, ""
	}

	return 200, res.RedirectURL
}
