package delivery

import (
	"altaproject2/config"
	"altaproject2/domain"
	"altaproject2/features/cart/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteCart(e *echo.Echo, ch domain.CartHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.PUT("/carts", ch.Update(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.POST("/carts", ch.Post(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/carts", ch.Get(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
