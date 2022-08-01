package factory

import (
	"altaproject2/features/user/delivery"
	"altaproject2/features/user/usecase"

	"altaproject2/features/user/data"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := data.New(db)
	valid := validator.New()
	userCase := usecase.New(userData, valid)
	userHandler := delivery.New(userCase)
	delivery.RouteUser(e, userHandler)

}
