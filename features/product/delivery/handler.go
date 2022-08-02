package delivery

import (
	"altaproject2/domain"
	"altaproject2/features/common"
	awss3 "altaproject2/utils/aws"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUseCase domain.ProductUseCase
	conn           *session.Session
}

func New(ps domain.ProductUseCase, aws *session.Session) domain.ProductHandler {
	return &productHandler{
		productUseCase: ps,
		conn:           aws,
	}
}

func (ps *productHandler) PostItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newproduct ProductFormat
		bind := c.Bind(&newproduct)
		//cost := 10

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		file, err := c.FormFile("photoprofile")

		if err != nil {
			log.Println(err)
		}

		filename := fmt.Sprintf("%s_profilepic.jpg", newproduct.ProductName)
		log.Println(filename)
		link := awss3.DoUpload(ps.conn, *file, filename)
		newproduct.ProductPic = link

		status := ps.productUseCase.PostItemAdmin(newproduct.ToModel())

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

func (ps *productHandler) UpdateItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newproduct ProductFormat
		// cost := 10
		param := common.ExtractData(c)
		bind := c.Bind(&newproduct)

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
			filename := fmt.Sprintf("%s_profilepic.jpg", newproduct.ProductName)
			log.Println(filename)
			link := awss3.DoUpload(ps.conn, *file, filename)
			newproduct.ProductPic = link
		}

		status := ps.productUseCase.UpdateItemAdmin(newproduct.ToModel(), param.ID)

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
