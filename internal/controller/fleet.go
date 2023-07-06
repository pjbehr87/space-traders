package controller

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/pjbehr87/space-traders/internal/service"
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

type shipContext struct {
	CanRefuel bool
	FuelPrice *int
}
type shipPage struct {
	Page PageData

	Ship        stapi.Ship
	Context     shipContext
	CurWp       *stapi.Waypoint
	SysWps      *[]stapi.Waypoint
	MarketGoods *[]stapi.MarketTradeGood
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
	e.POST("/my/ships/:shipSymbol/refuel", cont.refuelShip)
	e.POST("/my/ships/:shipSymbol/extract", cont.extractResources)
	e.POST("/my/ships/:shipSymbol/sell", cont.sellCargo)
}

// /////////////////////
// Ship Actions
// /////////////////////
func (ctl *fleetController) navigateShip(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	waypointSymbol := c.FormValue("waypointSymbol")
	navShipReq := *stapi.NewNavigateShipRequest(waypointSymbol)
	navigateShipResp, resp, err := ctl.sta.FleetApi.NavigateShip(c.Request().Context(), shipSymbol).NavigateShipRequest(navShipReq).Execute()
	if err != nil {
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}
		errMsgS := string(errMsg)
		c.Logger().Error(errMsgS)
		return c.String(http.StatusBadRequest, string(errMsgS))
	}

	err = c.JSON(http.StatusOK, navigateShipResp.Data)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
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
	purchaseShipResp, resp, err := ctl.sta.FleetApi.PurchaseShip(c.Request().Context()).PurchaseShipRequest(purchaseShipReq).Execute()
	if err != nil {
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}
		errMsgS := string(errMsg)
		c.Logger().Error(errMsgS)
		return c.String(http.StatusBadRequest, string(errMsgS))
	}

	err = c.JSON(http.StatusOK, purchaseShipResp.Data)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) orbitShip(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	orbitShipResp, resp, err := ctl.sta.FleetApi.OrbitShip(c.Request().Context(), shipSymbol).Execute()
	if err != nil {
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}
		errMsgS := string(errMsg)
		c.Logger().Error(errMsgS)
		return c.String(http.StatusBadRequest, string(errMsgS))
	}

	err = c.JSON(http.StatusOK, orbitShipResp.Data)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) dockShip(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	dockShipResp, resp, err := ctl.sta.FleetApi.DockShip(c.Request().Context(), shipSymbol).Execute()
	if err != nil {
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}
		errMsgS := string(errMsg)
		c.Logger().Error(errMsgS)
		return c.String(http.StatusBadRequest, string(errMsgS))
	}

	err = c.JSON(http.StatusOK, dockShipResp.Data)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) refuelShip(c echo.Context) error {
	fuelUnits := c.FormValue("units")
	fuelUnitsI, err := strconv.Atoi(fuelUnits)
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	refuelShipRequest := *stapi.NewRefuelShipRequest()
	if fuelUnitsI != 0 {
		refuelShipRequest.SetUnits(int32(fuelUnitsI))
	}
	shipSymbol := c.Param("shipSymbol")
	refuelShipResp, resp, err := ctl.sta.FleetApi.RefuelShip(c.Request().Context(), shipSymbol).RefuelShipRequest(refuelShipRequest).Execute()
	if err != nil {
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}
		errMsgS := string(errMsg)
		c.Logger().Error(errMsgS)
		return c.String(http.StatusBadRequest, string(errMsgS))
	}
	c.Logger().Debug("Refuel\n%+v", refuelShipResp)

	err = c.JSON(http.StatusOK, refuelShipResp.Data)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) extractResources(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	extractResourcesRequest := *stapi.NewExtractResourcesRequest()
	extractResourceResp, resp, err := ctl.sta.FleetApi.ExtractResources(c.Request().Context(), shipSymbol).ExtractResourcesRequest(extractResourcesRequest).Execute()
	if err != nil {
		errMsg, err := io.ReadAll(resp.Body)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusBadRequest, err.Error())
		}
		errMsgS := string(errMsg)
		c.Logger().Error(errMsgS)
		return c.String(http.StatusBadRequest, errMsgS)
	}
	c.Logger().Debug(fmt.Sprintf("ExtractResources\n%+v", extractResourceResp.Data))

	err = c.JSON(http.StatusOK, extractResourceResp.Data)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
func (ctl *fleetController) sellCargo(c echo.Context) error {
	shipSymbol := c.Param("shipSymbol")
	goodsSymbol := c.FormValue("symbol")
	goodsUnits, err := strconv.ParseInt(c.FormValue("units"), 10, 32)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	goodsSymbolE, err := stapi.NewTradeSymbolFromValue(goodsSymbol)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	sellCargoRequest := *stapi.NewSellCargoRequest(*goodsSymbolE, int32(goodsUnits))
	sellCargoResp, resp, err := ctl.sta.FleetApi.SellCargo(c.Request().Context(), shipSymbol).SellCargoRequest(sellCargoRequest).Execute()
	if err != nil {
		errMsg, err := io.ReadAll(resp.Body)
		errMsgS := string(errMsg)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.String(http.StatusBadRequest, string(errMsgS))
		}
		c.Logger().Error(errMsgS)
		return c.String(http.StatusBadRequest, string(errMsgS))
	}
	c.Logger().Debug(fmt.Sprintf("SellCargo\n%+v", sellCargoResp.Data))

	err = c.JSON(http.StatusOK, sellCargoResp.Data)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

// /////////////////////
// Fleet Management
// /////////////////////
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
	shipResp, _, err := ctl.sta.FleetApi.GetMyShip(c.Request().Context(), shipSymbol).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}
	ship := shipResp.Data
	wpSymbol := shipResp.Data.Nav.WaypointSymbol
	sysSymbol := shipResp.Data.Nav.SystemSymbol

	waypointsResp, _, err := ctl.sta.SystemsApi.GetSystemWaypoints(c.Request().Context(), sysSymbol).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}
	waypoints := waypointsResp.Data

	shipContext := shipContext{}
	var curWp *stapi.Waypoint
	var sysWps *[]stapi.Waypoint
	var marketGoods *[]stapi.MarketTradeGood
	if ship.Nav.Status != stapi.SHIPNAVSTATUS_IN_TRANSIT {
		// Capture the current WP and remove it from the "all waypoints" list
		for i, wp := range waypoints {
			if wp.Symbol == wpSymbol {
				curWp = &wp
				waypoints = append(waypoints[:i], waypoints[:i+1]...)
				break
			}
		}
		if ship.Nav.Status == stapi.SHIPNAVSTATUS_DOCKED && service.WpHasMarketplace(*curWp) {
			marketGoodsResp, _, err := ctl.sta.SystemsApi.GetMarket(c.Request().Context(), sysSymbol, wpSymbol).Execute()
			if err != nil {
				c.Logger().Error(err.Error())
				return c.String(http.StatusBadRequest, err.Error())
			}
			waypoints = nil
			marketGoods = &marketGoodsResp.Data.TradeGoods
			shipContext.CanRefuel, shipContext.FuelPrice = service.MarketHasFuel(marketGoodsResp.Data)
		}
		if ship.Nav.Status == stapi.SHIPNAVSTATUS_IN_ORBIT {
			sysWps = &waypoints
		}
	} else {
		sysWps = &waypoints
	}

	c.Logger().Debugf("Data: Ship\n%+v\nContext\n%+v\nCurrentWp\n%+v\nWaypoints\n%+v\nMarketGoods\n%+v", ship, shipContext, curWp, sysWps, marketGoods)

	err = c.Render(http.StatusOK, "ship", shipPage{
		Page: PageData{
			PageName: "Ship",
			JavaScript: []string{
				"ship",
			},
		},
		Ship:        ship,
		Context:     shipContext,
		CurWp:       curWp,
		SysWps:      sysWps,
		MarketGoods: marketGoods,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
