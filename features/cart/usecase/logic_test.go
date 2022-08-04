package usecase

import (
	"altaproject2/domain"
	"altaproject2/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestDeleteCart(t *testing.T) {
	repo := new(mocks.CartData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything).Return(true).Once()
		cartcase := New(repo, validator.New())
		delete, err := cartcase.DeleteCart(1, 1)

		assert.Nil(t, err)
		assert.Equal(t, true, delete)
		repo.AssertExpectations(t)
	})
}

func TestCreateCart(t *testing.T) {
	repo := new(mocks.CartData)
	mockData := domain.Cart{Userid: 1, Subtotal: 50000, Quantity: 2, Productid: 1}

	returnData := mockData
	returnData.ID = 1

	invalidData := mockData
	invalidData.Productid = 0

	t.Run("Success create", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("PostData", mock.Anything).Return(returnData).Once()
		cartcase := New(repo, validator.New())
		status := cartcase.PostCart(mockData)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Validator error", func(t *testing.T) {

		cartcase := New(repo, validator.New())
		status := cartcase.PostCart(invalidData)

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicated data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		cartcase := New(repo, validator.New())
		status := cartcase.PostCart(mockData)

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("Empty Data", func(t *testing.T) {
		returnData.ID = 0
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("PostData", mock.Anything).Return(returnData).Once()
		cartcase := New(repo, validator.New())
		status := cartcase.PostCart(mockData)

		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})
}

func TestUpdateCart(t *testing.T) {
	repo := new(mocks.CartData)
	mockData := domain.Cart{Userid: 1, Subtotal: 50000, Quantity: 2, Productid: 1}

	returnData := mockData
	returnData.ID = 1

	invalidData := mockData
	invalidData.ID = 0

	t.Run("Success update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(returnData).Once()
		cartcase := New(repo, validator.New())
		status := cartcase.UpdateData(mockData, 1)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("no productid", func(t *testing.T) {
		cartcase := New(repo, validator.New())
		status := cartcase.UpdateData(mockData, 0)

		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(invalidData).Once()
		cartcase := New(repo, validator.New())
		status := cartcase.UpdateData(mockData, 1)

		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})
}
