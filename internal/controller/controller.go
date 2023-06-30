package controller

import (
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type PageData struct {
	PageName   string
	JavaScript []string
	Css        []string
}

func InitRouter(e *echo.Echo, stl stlib.StLib) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	NewAgentController(e, stl)
	NewContractsController(e, stl)
	NewFleetController(e, stl)
	NewSystemsController(e, stl)
}
