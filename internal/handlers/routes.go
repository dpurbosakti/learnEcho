package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRoutes() *echo.Echo {
	// Config

	// Start routing
	r := echo.New()

	// logger endpoint
	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	r.GET("/home", Home)

	return r
}
