package delivery

import (
	"altaproject2/domain"

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
func (*orderHandler) Get() echo.HandlerFunc {
	panic("unimplemented")
}

// Post implements domain.OrderHandler
func (*orderHandler) Post() echo.HandlerFunc {
	panic("unimplemented")
}
