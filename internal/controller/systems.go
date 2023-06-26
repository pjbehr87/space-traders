package controller

import (
	"fmt"
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type SystemsController struct {
	stl stlib.StLib
}

func NewSystemsController(e *echo.Echo, stl stlib.StLib) {
	cont := SystemsController{
		stl: stl,
	}

	e.Logger.Debug("Router added: Systems")
	e.GET("/systems/:systemSymbol/waypoints/:waypointSymbol ", cont.getSystemWaypoints)
}

func (ctl *SystemsController) getSystemWaypoints(c echo.Context) error {
	waypoints, err := ctl.stl.GetWaypoint(c.Param("systemSymbol"), c.Param("waypointSymbol"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("%+v\n", waypoints))
}
