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

func NewFleetController(e *echo.Echo, stl stlib.StLib) {
	cont := fleetController{
		stl: stl,
	}

	e.Logger.Debug("Router added: Fleet")
	e.POST("/my/ships", cont.purchaseShip)
	e.GET("/my/ships", cont.getShips)
	e.GET("/my/ships/:shipSymbol", cont.getShip)
}

func (ctl *fleetController) purchaseShip(c echo.Context) error {
	c.Logger().Info(fmt.Sprintf("Request: POST PurchaseShip shipType=%s waypointSymbol=%s", c.FormValue("shipType"), c.FormValue("waypointSymbol")))
	err := ctl.stl.PurchaseShip(c.FormValue("shipType"), c.FormValue("waypointSymbol"))
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

func (ctl *fleetController) getShips(c echo.Context) error {
	c.Logger().Info("Request: GET Ships")
	ships, err := ctl.stl.GetShips()
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
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("--Resp: Ship\n%+v", ship)

	// err = c.Render(http.StatusOK, "myShips", shipsPage{
	// 	Page: PageData{
	// 		PageName: "Ship",
	// 	},
	// 	Ships: []stlib.Ship{ship},
	// })
	err = c.JSONPretty(http.StatusOK, shipsPage{
		Page: PageData{
			PageName: "Ship",
		},
		Ships: []stlib.Ship{ship},
	}, "   ")
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
