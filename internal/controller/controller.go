package controller

import (
	"net/http"

	stapi "github.com/pjbehr87/space-traders/st-api"
	stjobs "github.com/pjbehr87/space-traders/st-jobs"

	"github.com/labstack/echo/v4"
)

type PageData struct {
	PageName   string
	JavaScript []string
	Css        []string
}

func InitRouter(e *echo.Echo, sta *stapi.APIClient, stj stjobs.JobRunner) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	NewAgentController(e, sta)
	NewContractsController(e, sta)
	NewFleetController(e, sta)
	NewSystemsController(e, sta)
}
