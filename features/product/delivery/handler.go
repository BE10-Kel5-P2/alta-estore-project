package delivery

import (
	"altaproject2/domain"
	"altaproject2/features/common"
	awss3 "altaproject2/utils/aws"
	"fmt"
	"log"

	"net/http"
	"strconv"

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

		file, err := c.FormFile("productpic")

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

// DeleteItem implements domain.ProductHandler
func (ph *productHandler) DeleteItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := common.ExtractData(c)

		status, err := ph.productUserCase.DeleteItemAdmin(data.ID)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}

		if !status {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "Internal server error",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Success delete product",
		})
	}
}

// GetItem implements domain.ProductHandler
func (ph *productHandler) GetItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		prd := c.Param("prd")
		id, _ := strconv.Atoi(prd)

		data, err := ph.productUserCase.GetItemUser(id)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"productphoto": data.ProductPic,
			"productname":  data.ProductName,
			"price":        data.Price,
			"stock":        data.Stock,
			"quantity":     data.Qty,
			"description":  data.Description,
			"code":         200,
			"message":      "success get data",
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

// GetItems implements domain.ProductHandler
func (ps *productHandler) GetItems() echo.HandlerFunc {
	return func(c echo.Context) error {
		var arrmap []map[string]interface{}
		data, status := ps.productUseCase.GetAllItems()

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}
		for i := 0; i < len(data); i++ {
			var res = map[string]interface{}{}
			res["productphoto"] = data[0].ProductPic
			res["productname"] = data[0].ProductName
			res["price"] = data[0].Price
			res["stock"] = data[0].Stock
			res["qty"] = data[0].Qty
			arrmap = append(arrmap, res)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    arrmap,
			"code":    status,
			"message": "Data not found",
		})
	}
}
