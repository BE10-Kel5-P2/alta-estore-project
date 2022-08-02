package delivery

import (
	"altaproject2/config"
	"altaproject2/domain"
	"altaproject2/features/product/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteProduct(e *echo.Echo, ph domain.ProductHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/products", ph.GetItems())
	e.GET("/products/:id", ph.GetItems())
  e.POST("/products", psr.PostItem(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/products", psr.UpdateItem(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
