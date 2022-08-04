package delivery

import (
	"altaproject2/config"
	"altaproject2/domain"
	"altaproject2/features/order/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteOrder(e *echo.Echo, oh domain.OrderHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/orders", oh.Post(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
