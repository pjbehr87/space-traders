package controller

import (
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type systemsController struct {
	stl stlib.StLib
}

type shipyardParams struct {
	Waypoint string
}

type waypointsPage struct {
	Page PageData

	Waypoints []stlib.Waypoint
}
type shipyardPage struct {
	Page PageData

	Shipyard stlib.Shipyard
	Params   shipyardParams
}

func NewSystemsController(e *echo.Echo, stl stlib.StLib) {
	cont := systemsController{
		stl: stl,
	}

	e.Logger.Debug("Router added: Systems")
	e.GET("/systems/:systemSymbol/waypoints", cont.listSystemWaypoints)
	e.GET("/systems/:systemSymbol/waypoints/:waypointSymbol", cont.getSystemWaypoints)
	e.GET("/systems/:systemSymbol/waypoints/:waypointSymbol/shipyard", cont.getShipyard)
}

func (ctl *systemsController) listSystemWaypoints(c echo.Context) error {
	c.Logger().Info("Request: LIST Waypoints")
	waypoints, err := ctl.stl.ListWaypoints(c.Param("systemSymbol"))
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Resp: Waypoints\n%+v", waypoints)

	err = c.Render(http.StatusOK, "waypoints", waypointsPage{
		Page: PageData{
			PageName: "Waypoints",
		},
		Waypoints: waypoints,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *systemsController) getSystemWaypoints(c echo.Context) error {
	c.Logger().Info("Request: GET Waypoint")
	waypoint, err := ctl.stl.GetWaypoint(c.Param("systemSymbol"), c.Param("waypointSymbol"))
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Resp: Waypoint\n%+v", waypoint)

	// We can share the get waypoint and list waypoints page. Just pass as single-item slice
	err = c.Render(http.StatusOK, "waypoints", waypointsPage{
		Page: PageData{
			PageName: "Waypoint",
		},
		Waypoints: []stlib.Waypoint{waypoint},
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *systemsController) getShipyard(c echo.Context) error {
	c.Logger().Info("Request: GET Shipyard")
	shipyard, err := ctl.stl.GetShipyard(c.Param("systemSymbol"), c.Param("waypointSymbol"))
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Resp: Waypoint\n%+v", shipyard)

	err = c.Render(http.StatusOK, "shipyard", shipyardPage{
		Page: PageData{
			PageName: "Waypoint",
			JavaScript: []string{
				"shipyard",
			},
		},
		Params: shipyardParams{
			Waypoint: c.Param("waypointSymbol"),
		},
		Shipyard: shipyard,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
