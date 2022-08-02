package delivery

import (
	"altaproject2/domain"
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
func (*productHandler) GetItem() echo.HandlerFunc {
	panic("unimplemented")
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
