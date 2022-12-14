package delivery

import (
	"altaproject2/domain"
	common "altaproject2/features/common"
	awss3 "altaproject2/utils/aws"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCase
	conn        *session.Session
}

func New(us domain.UserUseCase, aws *session.Session) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
		conn:        aws,
	}
}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		var arrmap = map[string]interface{}{}

		errLog := c.Bind(&userLogin)

		if errLog != nil {
			log.Println("invalid input")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		data, err := us.userUsecase.Login(userLogin.ToModelLogin())

		if err != nil {
			log.Println("Login failed", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Wrong username or password",
			})
		}

		token := common.GenerateToken(data)

		arrmap["token"] = token
		arrmap["username"] = data.Username
		arrmap["role"] = data.Role

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    arrmap,
			"code":    200,
			"message": "Login success",
		})
	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := common.ExtractData(c)

		status, err := us.userUsecase.Delete(data.ID)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		if !status {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success delete user",
		})
	}
}

// Register implements domain.UserHandler
func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat
		bind := c.Bind(&newuser)
		cost := 10

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := us.userUsecase.RegisterUser(newuser.ToModel(), cost)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Register success",
		})
	}
}

// Update implements domain.UserHandler
func (us *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newuser UserFormat
		cost := 10
		param := common.ExtractData(c)
		bind := c.Bind(&newuser)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		file, err := c.FormFile("photoprofile")

		if err == nil {
			log.Println(err)
			filename := fmt.Sprintf("%s_profilepic.jpg", newuser.Username)
			log.Println(filename)
			link := awss3.DoUpload(us.conn, *file, filename)
			newuser.Photoprofile = link
		}

		status := us.userUsecase.UpdateUser(newuser.ToModel(), param.ID, cost)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Update success",
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var arrmap = []map[string]interface{}{}
		usr := common.ExtractData(c)

		data, datauc, err := uh.userUsecase.GetProfile(usr.ID)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		for i := 0; i < len(datauc); i++ {
			var res = map[string]interface{}{}
			res["productname"] = datauc[i].ProductName
			res["quantity"] = datauc[i].Quantity
			res["totalprice"] = datauc[i].Subtotal

			arrmap = append(arrmap, res)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"username":     data.Username,
			"photoprofile": data.PhotoProfile,
			"data":         arrmap,
			"code":         200,
			"message":      "get data success",
		})
	}
}
