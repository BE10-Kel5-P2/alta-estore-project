package usecase

import (
	"altaproject2/domain"
	"altaproject2/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProduct(t *testing.T) {
	repo := new(mocks.CartData)
	mockData := domain.Cart{ID: 1, Userid: 1, Subtotal: 50000, Quantity: 2, Productid: 1}

	returnData := mockData
	// returnData.ID = 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateData", returnData, 1).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateData(mockData, 1)

		assert.Equal(t, 200, res)
	})
}
