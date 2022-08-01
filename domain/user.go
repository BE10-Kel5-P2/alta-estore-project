package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID           int
	Fullname     string
	Username     string
	Email        string
	Address      string
	PhotoProfile string
	Password     string
	Role         string
}

type UserHandler interface {
	Login() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type UserUseCase interface {
	Login(Email string, password string) (username string, role string, token string, err error)
	Delete(userID int) (bool, error)
}

type UserData interface {
	Login(email string, password string) (username string, role string, token string, err error)
	Delete(userID int) bool
}
