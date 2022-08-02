package factory

import (
	ud "altaproject2/features/user/data"
	udeli "altaproject2/features/user/delivery"
	uc "altaproject2/features/user/usecase"

	pd "altaproject2/features/product/data"
	pdeli "altaproject2/features/product/delivery"
	pc "altaproject2/features/product/usecase"

	cd "altaproject2/features/cart/data"
	cdeli "altaproject2/features/cart/delivery"
	cc "altaproject2/features/cart/usecase"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB, awsConn *session.Session) {
	valid := validator.New()

	userData := ud.New(db)
	userCase := uc.New(userData, valid)
	userHandler := udeli.New(userCase, awsConn)
	udeli.RouteUser(e, userHandler)

	postData := pd.New(db)
	postCase := pc.New(postData, valid)
	postHandler := pdeli.New(postCase, awsConn)
	pdeli.RouteProduct(e, postHandler)

	cartData := cd.New(db)
	cartCase := cc.New(cartData)
	cartHandler := cdeli.New(cartCase)
	cdeli.RouteCart(e, cartHandler)
}
