package controller

import (
	"fmt"
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type fleetController struct {
	stl stlib.StLib
}

type shipsPage struct {
	Page PageData

	Ships []stlib.Ship
}
type shipPage struct {
	Page PageData

	Ship      stlib.Ship
	Waypoints []stlib.Waypoint
}

func NewFleetController(e *echo.Echo, stl stlib.StLib) {
	cont := fleetController{
		stl: stl,
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
	c.Logger().Info(fmt.Sprintf("Request: POST PurchaseShip shipType=%s waypointSymbol=%s", c.FormValue("shipType"), c.FormValue("waypointSymbol")))
	err := ctl.stl.PurchaseShip(c.FormValue("shipType"), c.FormValue("waypointSymbol"))
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
func (ctl *fleetController) orbitShip(c echo.Context) error {
	c.Logger().Info(fmt.Sprintf("Request: POST OrbitShip shipSymbol=%s", c.Param("shipSymbol")))
	err := ctl.stl.OrbitShip(c.Param("shipSymbol"))
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
	c.Logger().Info(fmt.Sprintf("Request: POST DockShip shipSymbol=%s", c.Param("shipSymbol")))
	err := ctl.stl.DockShip(c.Param("shipSymbol"))
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
	ships, err := ctl.stl.ListShips()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Resp: Ships\n%+v", ships)

	err = c.Render(http.StatusOK, "myShips", shipsPage{
		Page: PageData{
			PageName: "Ships",
		},
		Ships: ships,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) getShip(c echo.Context) error {
	c.Logger().Info("Request: GET Ship")
	ship, err := ctl.stl.GetShip(c.Param("shipSymbol"))
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}

	waypoints, err := ctl.stl.ListWaypoints(ship.Nav.WaypointSymbol.System)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Logger().Debugf("Data: Ship\n%+v\nWaypoints\n%+v", ship, waypoints)
	err = c.Render(http.StatusOK, "ship", shipPage{
		Page: PageData{
			PageName: "Ship",
			JavaScript: []string{
				"ship",
			},
		},
		Ship:      ship,
		Waypoints: waypoints,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *fleetController) navigateShip(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	waypointSymbol := c.FormValue("waypointSymbol")
	c.Logger().Info(fmt.Sprintf("Request: POST NavigateShip shipSymbol=%s waypointSymbol=%s", shipSymbol, waypointSymbol))
	err := ctl.stl.NavigateShip(shipSymbol, waypointSymbol)
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
