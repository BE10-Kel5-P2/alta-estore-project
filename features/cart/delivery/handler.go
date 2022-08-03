package delivery

import (
	"altaproject2/domain"
	"altaproject2/features/common"
<<<<<<< HEAD
=======
	"log"
>>>>>>> 025aa539018e15a8e76492712446454f0bbd745e
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
func (cs *cartHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newcart CartFormat
		token := common.ExtractData(c)
		newcart.Userid = token.ID
		bind := c.Bind(&newcart)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := cs.cartUseCase.PostCart(newcart.ToModel())

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
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
			"message": "Success add item",
		})
	}
}

// Update implements domain.CartHandler
func (ch *cartHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updatecart CartFormat
		bind := c.Bind(&updatecart)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := ch.cartUseCase.UpdateData(updatecart.ToModel(), updatecart.Productid)

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
