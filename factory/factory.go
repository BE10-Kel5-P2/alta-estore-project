package factory

import (
	"altaproject2/features/user/data"
	"altaproject2/features/user/delivery"
	"altaproject2/features/user/usecase"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB, awsConn *session.Session) {
	userData := data.New(db)
	valid := validator.New()
	userCase := usecase.New(userData, valid)
	userHandler := delivery.New(userCase, awsConn)
	delivery.RouteUser(e, userHandler)

}
