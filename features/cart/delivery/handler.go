package delivery

import (
	"altaproject2/domain"
	"altaproject2/features/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	cartUseCase domain.CartUseCase
}

func New(cs domain.CartUseCase) domain.CartHandler {
	return &cartHandler{
		cartUseCase: cs,
	}
}

// Delete implements domain.CartHandler
func (cs *cartHandler) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := common.ExtractData(c)

		status, err := cs.cartUseCase.DeleteCart(data.ID)

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
			"message": "success delete product",
		})
	}
}

// Get implements domain.CartHandler
func (*cartHandler) Get() echo.HandlerFunc {
	panic("unimplemented")
}

// Post implements domain.CartHandler
func (*cartHandler) Post() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements domain.CartHandler
func (*cartHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}
