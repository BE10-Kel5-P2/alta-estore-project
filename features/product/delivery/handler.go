package delivery

import (
	"altaproject2/domain"

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
func (*productHandler) GetItems() echo.HandlerFunc {
	panic("unimplemented")
}

// PostCart implements domain.ProductHandler
func (*productHandler) PostCart() echo.HandlerFunc {
	panic("unimplemented")
}

// PostItem implements domain.ProductHandler
func (*productHandler) PostItem() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdateItem implements domain.ProductHandler
func (*productHandler) UpdateItem() echo.HandlerFunc {
	panic("unimplemented")
}
