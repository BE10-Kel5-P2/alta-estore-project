package usecase

import (
	"altaproject2/domain"
	"altaproject2/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUser(t *testing.T) {
	repo := new(mocks.ProductData)

	mockData := []domain.Product{{ProductName: "Men's bag", Description: "ini adalah tas pria",
		Price: 200000, ProductPic: "bag.jpg", Stock: 5, Qty: 2}, {ProductName: "Men's bag", Description: "ini adalah tas pria",
		Price: 200000, ProductPic: "bag.jpg", Stock: 5, Qty: 2}}

	t.Run("Success get all data", func(t *testing.T) {
		repo.On("GetAllItemData", mock.Anything).Return(mockData).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetAllItems()

		assert.Equal(t, 200, status)
		assert.Equal(t, mockData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetAllItemData", mock.Anything).Return(nil).Once()
		useCase := New(repo, validator.New())
		res, status := useCase.GetAllItems()

		assert.Equal(t, 404, status)
		assert.Equal(t, []domain.Product(nil), res)
		repo.AssertExpectations(t)
	})
}

func TestPostProduct(t *testing.T) {
	repo := new(mocks.ProductData)

	mockData := domain.Product{ProductName: "Men's bag", Description: "ini adalah tas pria",
		Price: 200000, ProductPic: "bag.jpg", Stock: 5, Qty: 2}
	// emptyData := domain.Product{ProductName: "", Description: "",
	// Price: 0, ProductPic: "", Stock: 0, Qty:0 }

	returnData := mockData
	returnData.ID = 1

	invalidData := mockData
	invalidData.ProductName = ""

	noData := mockData
	noData.ID = 0

	t.Run("Success insert data", func(t *testing.T) {
		// useCase := New(&mockUserDataTrue{})
		repo.On("PostItemData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.PostItemAdmin(mockData)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	mockData := domain.Product{ProductName: "Men's bag", Description: "ini adalah tas pria",
		Price: 200000, ProductPic: "bag.jpg", Stock: 5, Qty: 2}

	returnData := mockData
	returnData.ID = 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateItemData", returnData, 1).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateItemAdmin(mockData, 1)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})
}
