package delivery

import (
	"altaproject2/domain"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	cartUseCase domain.CartUseCase
}

func New(ps domain.CartUseCase) domain.CartHandler {
	return &cartHandler{
		cartUseCase: ps,
	}
}

// Delete implements domain.CartHandler
func (*cartHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
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
