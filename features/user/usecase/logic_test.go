package usecase

import (
	"altaproject2/domain"
	"altaproject2/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertUser(t *testing.T) {
	repo := new(mocks.UserData)
	cost := 10

	mockData := domain.User{Username: "NotAPanda", Email: "lukman@gmail.com", Address: "jakarta", Password: "polar", PhotoProfile: "lukman.jpg"}
	emptyData := domain.User{Username: "", Email: "", Address: "", Password: "", PhotoProfile: ""}

	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru"

	invalidData := mockData
	invalidData.Email = ""

	noData := mockData
	noData.ID = 0

	t.Run("Success insert data", func(t *testing.T) {
		// useCase := New(&mockUserDataTrue{})
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Validator Error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(invalidData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, 40)

		assert.Equal(t, 500, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(emptyData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(noData, cost)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)
	cost := 10

	mockData := domain.User{Username: "NotAPanda", Email: "lukman@gmail.com", Address: "jakarta", Password: "polar", PhotoProfile: "lukman.jpg"}

	returnData := mockData
	returnData.ID = 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("UpdateUserData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, cost)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 0, cost)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, 40)

		assert.Equal(t, 500, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{Username: "NotAPanda", Password: "polar"}
	returnData := domain.User{ID: 1}
	token := "sjage2w62vsdgaqsgh"
	t.Run("Succes Login", func(t *testing.T) {
		repo.On("GetPasswordData", mock.Anything).Return("$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru")
		repo.On("Login", mock.Anything).Return(returnData, token).Once()
		userUseCase := New(repo, validator.New())
		res, err := userUseCase.Login(mockData)

		assert.Nil(t, err)
		assert.Greater(t, res.ID, 0)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New())
		delete, err := usecase.Delete(1)

		assert.Nil(t, err)
		assert.Equal(t, true, delete)
		repo.AssertExpectations(t)
	})
}

func TestGetUser(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{ID: 1, Username: "NotAPanda", Email: "lukman@gmail.com", Address: "jakarta", Password: "polar", PhotoProfile: "lukman.jpg"}
	returnusercart := []domain.UserCart{{ProductName: "men bag", Quantity: 1, Subtotal: 100000}, {ProductName: "men shoes", Quantity: 1, Subtotal: 100000}}
	t.Run("success get data", func(t *testing.T) {
		repo.On("GetUserCartData", mock.Anything).Return(returnusercart).Once()
		repo.On("GetProfile", mock.Anything).Return(mockData, nil).Once()
		useCase := New(repo, validator.New())
		res, uc, error := useCase.GetProfile(1)

		assert.Nil(t, error)
		assert.Equal(t, returnusercart, uc)
		assert.Equal(t, mockData.Username, res.Username)
		assert.Equal(t, mockData.PhotoProfile, res.PhotoProfile)
		repo.AssertExpectations(t)
	})
}
