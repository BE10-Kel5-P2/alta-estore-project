package usecase

import (
	"errors"

	"altaproject2/domain"

	"github.com/go-playground/validator/v10"
)

type userCase struct {
	userData domain.UserData
	valid    *validator.Validate
}

func New(ud domain.UserData, val *validator.Validate) domain.UserUseCase {
	return &userCase{
		userData: ud,
		valid:    val,
	}
}

func (ud *userCase) Login(email string, password string) (username string, role string, token string, err error) {
	username, role, token, err = ud.userData.Login(email, password)
	return username, role, token, err
}

func (ud *userCase) Delete(userId int) (bool, error) {
	res := ud.userData.Delete(userId)

	if !res {
		return false, errors.New("failed to delete user")
	}
	return true, nil
}
