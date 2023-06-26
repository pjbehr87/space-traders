package controller

import (
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type SystemsController struct {
	stl stlib.StLib
}

type WaypointPage struct {
	Page

	Waypoints []stlib.Waypoint
}

func NewSystemsController(e *echo.Echo, stl stlib.StLib) {
	cont := SystemsController{
		stl: stl,
	}

	e.Logger.Debug("Router added: Systems")
	e.GET("/systems/:systemSymbol/waypoints/:waypointSymbol ", cont.getSystemWaypoints)
}

func (ctl *SystemsController) getSystemWaypoints(c echo.Context) error {
	c.Logger().Info("GET Waypoint")
	waypoints, err := ctl.stl.GetWaypoint(c.Param("systemSymbol"), c.Param("waypointSymbol"))
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Waypoints: %+v", waypoints)

	err = c.Render(http.StatusOK, "waypoint", WaypointPage{
		Page: Page{
			PageName: "Waypoint",
		},
		Waypoints: waypoints,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
