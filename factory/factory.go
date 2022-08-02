package factory

import (
	"altaproject2/features/user/data"
	"altaproject2/features/user/delivery"
	"altaproject2/features/user/usecase"

	pd "altaproject2/features/product/data"
	pDelivery "altaproject2/features/product/delivery"
	productCases "altaproject2/features/product/usecase"

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

	productData := pd.New(db)
	valid = validator.New()
	productCase := productCases.New(productData, valid)
	productHandler := pDelivery.New(productCase, awsConn)
	pDelivery.RouteProduct(e, productHandler)

}
