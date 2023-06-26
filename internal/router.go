package internal

import (
	"net/http"
	"space-traders/internal/controller"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, stl stlib.StLib) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	controller.NewAgentController(e, stl)
	controller.NewSystemsController(e, stl)
}
