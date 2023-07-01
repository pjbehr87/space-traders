package controller

import (
	"io"
	"net/http"

	stapi "github.com/pjbehr87/space-traders/st-api"

	"github.com/labstack/echo/v4"
)

type fleetController struct {
	sta *stapi.APIClient
}

type shipsPage struct {
	Page PageData

	Ships []stapi.Ship
}
type shipPage struct {
	Page PageData

	Ship      stapi.Ship
	Waypoints []stapi.Waypoint
}

func NewFleetController(e *echo.Echo, sta *stapi.APIClient) {
	cont := fleetController{
		sta: sta,
	}

	e.Logger.Debug("Router added: Fleet")
	e.POST("/my/ships", cont.purchaseShip)
	e.GET("/my/ships", cont.listShips)
	e.GET("/my/ships/:shipSymbol", cont.getShip)
	e.POST("/my/ships/:shipSymbol/orbit", cont.orbitShip)
	e.POST("/my/ships/:shipSymbol/dock", cont.dockShip)
	e.POST("/my/ships/:shipSymbol/navigate", cont.navigateShip)
}

func (ctl *fleetController) purchaseShip(c echo.Context) error {
	shipType, err := stapi.NewShipTypeFromValue(c.FormValue("shipType"))
	if err != nil {
		c.Logger().Error(err)
		c.String(http.StatusBadRequest, err.Error())
		return err
	}
	waypointSymbol := c.FormValue("waypointSymbol")

	c.Logger().Debug(*shipType, waypointSymbol)

	purchaseShipReq := *stapi.NewPurchaseShipRequest(*shipType, waypointSymbol)
	_, resp, err := ctl.sta.FleetApi.PurchaseShip(c.Request().Context()).PurchaseShipRequest(purchaseShipReq).Execute()
	if err != nil {
		errBody, _ := io.ReadAll(resp.Body)
		c.Logger().Error(string(errBody))
		return c.String(http.StatusBadRequest, string(errBody))
	}

	err = c.String(http.StatusOK, "{}")
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) orbitShip(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	_, _, err := ctl.sta.FleetApi.OrbitShip(c.Request().Context(), shipSymbol).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = c.String(http.StatusOK, "{}")
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) dockShip(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	_, _, err := ctl.sta.FleetApi.DockShip(c.Request().Context(), shipSymbol).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = c.String(http.StatusOK, "{}")
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *fleetController) listShips(c echo.Context) error {
	c.Logger().Info("Request: LIST Ships")

	myShips, _, err := ctl.sta.FleetApi.GetMyShips(c.Request().Context()).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Resp: Ships\n%+v", myShips.Data)

	err = c.Render(http.StatusOK, "myShips", shipsPage{
		Page: PageData{
			PageName: "Ships",
		},
		Ships: myShips.Data,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) getShip(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	ship, _, err := ctl.sta.FleetApi.GetMyShip(c.Request().Context(), shipSymbol).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}

	systemSymbol := ship.Data.Nav.SystemSymbol
	waypoints, _, err := ctl.sta.SystemsApi.GetSystemWaypoints(c.Request().Context(), systemSymbol).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Logger().Debugf("Data: Ship\n%+v\nWaypoints\n%+v", ship.Data, waypoints.Data)
	err = c.Render(http.StatusOK, "ship", shipPage{
		Page: PageData{
			PageName: "Ship",
			JavaScript: []string{
				"ship",
			},
		},
		Ship:      ship.Data,
		Waypoints: waypoints.Data,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *fleetController) navigateShip(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	waypointSymbol := c.FormValue("waypointSymbol")
	navShipReq := *stapi.NewNavigateShipRequest(waypointSymbol)
	_, _, err := ctl.sta.FleetApi.NavigateShip(c.Request().Context(), shipSymbol).NavigateShipRequest(navShipReq).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = c.NoContent(http.StatusOK)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
