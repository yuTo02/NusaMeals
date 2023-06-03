package route

import (
	"reglog/controller"
	m "reglog/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	m.LoggerMiddleware(e)

	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)
	e.POST("/logout", controller.LogoutController)

	return e
}
