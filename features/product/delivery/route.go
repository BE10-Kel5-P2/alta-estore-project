package delivery

import (
	"altaproject2/domain"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, usr domain.UserHandler) {
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	// 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	// }))
	// e.Pre(middleware.RemoveTrailingSlash())

}
