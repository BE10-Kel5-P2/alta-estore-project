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
