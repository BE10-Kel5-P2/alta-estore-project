package usecase

import (
	"altaproject2/domain"
	"altaproject2/domain/mocks"
	"testing"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostCart(t *testing.T) {
	repo := new(mocks.OrderData)

	mockOrder := domain.Order{Userid: 1, Total: 400000, Data: []domain.ProductOrders{}}

	mockProductOrders := []domain.ProductOrders{{Orderid: 1, Productid: 1, Price: 100000, Quantity: 2, Subtotal: 200000, Status: 0, Payment: "payment"},
		{Orderid: 1, Productid: 2, Price: 100000, Quantity: 2, Subtotal: 200000, Status: 0, Payment: "payment"}}

	returnData := []domain.ProductOrders{{ID: 1, Orderid: 1, Productid: 1, Price: 100000, Quantity: 2, Subtotal: 200000, Status: 0, Payment: "payment"}, {ID: 2, Orderid: 1, Productid: 2, Price: 100000, Quantity: 2, Subtotal: 200000, Status: 0, Payment: "payment"}}

	t.Run("success post cart", func(t *testing.T) {
		repo.On("PostOrderData", mock.Anything).Return(1).Once()
		repo.On("PostProductOrderData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, snap.Client{ServerKey: "SB-Mid-server-RdSR7SzPE67DEllONzb7sCOX", Env: midtrans.Environment})
		status, url := useCase.PostOrder(mockOrder, mockProductOrders)

		assert.NotNil(t, url)
		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("PostOrderData", mock.Anything).Return(0).Once()
		useCase := New(repo, snap.Client{})
		status, _ := useCase.PostOrder(mockOrder, mockProductOrders)

		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})
}
