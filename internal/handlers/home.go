package handlers

import (
	"net/http"

	t "learn-echo/pkg/types"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	data := t.II{
		"message": "Welcome !!",
	}
	return c.JSON(http.StatusOK, data)
}
