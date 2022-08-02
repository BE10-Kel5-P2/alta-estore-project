package delivery

import (
	"altaproject2/domain"
	"altaproject2/features/common"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUserCase domain.ProductUseCase
	conn            *session.Session
}

func New(ps domain.ProductUseCase, aws *session.Session) domain.ProductHandler {
	return &productHandler{
		productUserCase: ps,
		conn:            aws,
	}
}

// DeleteItem implements domain.ProductHandler
func (*productHandler) DeleteItem() echo.HandlerFunc {
	panic("unimplemented")
}

// GetItem implements domain.ProductHandler
func (ph *productHandler) GetItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		prd := common.ExtractData(c)

		data, err := ph.productUserCase.GetItemUser(prd.ID)

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

// GetItems implements domain.ProductHandler
func (ps *productHandler) GetItems() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, status := ps.productUserCase.GetAllItems()

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"code":    status,
			"message": "Data not found",
		})
	}
}

// PostItem implements domain.ProductHandler
func (*productHandler) PostItem() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdateItem implements domain.ProductHandler
func (*productHandler) UpdateItem() echo.HandlerFunc {
	panic("unimplemented")
}
