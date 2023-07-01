package controller

import (
	"net/http"

	stapi "github.com/pjbehr87/space-traders/st-api"

	"github.com/labstack/echo/v4"
)

type systemsController struct {
	sta *stapi.APIClient
}

type shipyardParams struct {
	Waypoint string
}

type waypointsPage struct {
	Page PageData

	Waypoints []stapi.Waypoint
}
type shipyardPage struct {
	Page PageData

	Shipyard stapi.Shipyard
	Params   shipyardParams
}

func NewSystemsController(e *echo.Echo, sta *stapi.APIClient) {
	cont := systemsController{
		sta: sta,
	}

	e.Logger.Debug("Router added: Systems")
	e.GET("/systems/:systemSymbol/waypoints", cont.listSystemWaypoints)
	e.GET("/systems/:systemSymbol/waypoints/:waypointSymbol", cont.getSystemWaypoints)
	e.GET("/systems/:systemSymbol/waypoints/:waypointSymbol/shipyard", cont.getShipyard)
}

func (ctl *systemsController) listSystemWaypoints(c echo.Context) error {
	systemSymbol := c.Param("systemSymbol")
	waypoints, _, err := ctl.sta.SystemsApi.GetSystemWaypoints(c.Request().Context(), systemSymbol).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Resp: Waypoints\n%+v", waypoints)

	err = c.Render(http.StatusOK, "waypoints", waypointsPage{
		Page: PageData{
			PageName: "Waypoints",
		},
		Waypoints: waypoints.Data,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *systemsController) getSystemWaypoints(c echo.Context) error {
	systemSymbol := c.Param("systemSymbol")
	waypointSymbol := c.Param("waypointSymbol")
	waypoint, _, err := ctl.sta.SystemsApi.GetWaypoint(c.Request().Context(), systemSymbol, waypointSymbol).Execute()
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
		Waypoints: []stapi.Waypoint{waypoint.Data},
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *systemsController) getShipyard(c echo.Context) error {
	systemSymbol := c.Param("systemSymbol")
	waypointSymbol := c.Param("waypointSymbol")
	shipyard, _, err := ctl.sta.SystemsApi.GetShipyard(c.Request().Context(), systemSymbol, waypointSymbol).Execute()
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
			Waypoint: waypointSymbol,
		},
		Shipyard: shipyard.Data,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
