package usecase

import (
	"altaproject2/domain"
	"altaproject2/domain/mocks"
	"errors"
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
		repo.On("PostItemData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.PostItemAdmin(mockData)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Validation error", func(t *testing.T) {
		noData.Stock = 0
		useCase := New(repo, validator.New())
		res := useCase.PostItemAdmin(noData)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})

	t.Run("Empty data", func(t *testing.T) {
		returnData.ID = 0
		repo.On("PostItemData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.PostItemAdmin(mockData)

		assert.Equal(t, 404, res)
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

	t.Run("Data not found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.UpdateItemAdmin(mockData, 0)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})
}

func TestGetProduct(t *testing.T) {
	repo := new(mocks.ProductData)

	mockData := domain.Product{ID: 1, ProductName: "Men's Shirt", Description: "Ready for size M, L, XL, XXL", Price: 100000, ProductPic: "Shirt.jpg", Stock: 5, Qty: 1}

	t.Run("success get data", func(t *testing.T) {
		repo.On("GetItemData", mock.Anything).Return(mockData, nil).Once()
		productCase := New(repo, validator.New())
		res, error := productCase.GetItemUser(1)

		assert.Nil(t, error)
		assert.Equal(t, mockData.ProductName, res.ProductName)
		assert.Equal(t, mockData.Description, res.Description)
		assert.Equal(t, mockData.Price, res.Price)
		assert.Equal(t, mockData.Stock, res.Stock)
		assert.Equal(t, mockData.Qty, res.Qty)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetItemData", mock.Anything).Return(mockData, errors.New("Error")).Once()
		productCase := New(repo, validator.New())
		res, error := productCase.GetItemUser(0)

		assert.NotNil(t, error)
		assert.Equal(t, domain.Product{}, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	repo := new(mocks.ProductData)

	t.Run("Succes delete product", func(t *testing.T) {
		repo.On("DeleteItemData", mock.Anything).Return(true).Once()
		productcase := New(repo, validator.New())
		delete, err := productcase.DeleteItemAdmin(1)

		assert.Nil(t, err)
		assert.Equal(t, true, delete)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete data", func(t *testing.T) {
		repo.On("DeleteItemData", mock.Anything).Return(false).Once()
		productcase := New(repo, validator.New())
		delete, err := productcase.DeleteItemAdmin(1)

		assert.NotNil(t, err)
		assert.Equal(t, false, delete)
		repo.AssertExpectations(t)
	})
}
