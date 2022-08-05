package delivery

import (
	"altaproject2/domain"
	"altaproject2/features/common"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderUseCase domain.OrderUseCase
}

func New(oc domain.OrderUseCase) domain.OrderHandler {
	return &orderHandler{
		orderUseCase: oc,
	}
}

// Delete implements domain.OrderHandler
func (*orderHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// Get implements domain.OrderHandler
func (oc *orderHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		prd := c.Param("id")
		id, _ := strconv.Atoi(prd)

		data, err := oc.orderUseCase.GetOrder(id)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    404,
				"message": "Data not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"orderid":   data.Orderid,
			"productid": data.Productid,
			"price":     data.Price,
			"quantity":  data.Quantity,
			"subtotal":  data.Subtotal,
			"status":    data.Status,
			"payment":   data.Payment,
			"code":      200,
			"message":   "success get data",
		})
	}
}

// Post implements domain.OrderHandler
func (oc *orderHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data []domain.ProductOrders
		var neworder OrderFormat
		var newproductorder []ProductOrdersFormat

		token := common.ExtractData(c)
		neworder.Userid = token.ID
		neworder.Data = newproductorder

		bind := c.Bind(&neworder)

		if bind != nil {
			log.Println("cant bind", bind.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    400,
				"message": "Wrong input",
			})
		}

		total := oc.orderUseCase.Sum(neworder.ToModel())

		neworder.Total = total

		for _, val := range neworder.Data {
			data = append(data, val.ToModelProductOrders())
		}

		status, url := oc.orderUseCase.PostOrder(neworder.ToModel(), data)

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"paymentlink": url,
			"code":        status,
			"message":     "Success send data",
		})
	}

}
