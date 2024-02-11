package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterMiddleWare(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// csrf := middleware.DefaultCSRFConfig
	// csrf.TokenLookup = ":form:_csrf"
	// e.Use(middleware.CSRFWithConfig(csrf))
	e.Use(middleware.CSRF())

	e.Use(middleware.CORS())

}
