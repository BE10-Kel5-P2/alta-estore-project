package delivery

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"altaproject2/domain"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(us domain.UserUseCase) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
	}
}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		errLog := c.Bind(&userLogin)
		if errLog != nil {
			return c.JSON(http.StatusBadRequest, "email atau password salah")
		}
		username, role, token, err := us.userUsecase.Login(userLogin.Email, userLogin.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "login gagal")
		}
		data := map[string]interface{}{
			"username": username,
			"role":     role,
			"token":    token,
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "login berhasil",
			"data":    data,
		})
	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := us.userUsecase.Delete(cnv)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user",
		})
	}
}
