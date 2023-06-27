package controller

import (
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type ContractsController struct {
	stl stlib.StLib
}

type ContractsPage struct {
	Page

	Contracts []stlib.Contract
}

func NewContractsController(e *echo.Echo, stl stlib.StLib) {
	cont := ContractsController{
		stl: stl,
	}

	e.Logger.Debug("Router added: Contracts")
	e.GET("/my/contracts", cont.listContracts)
}

func (ctl *ContractsController) listContracts(c echo.Context) error {
	c.Logger().Info("GET Contracts")
	contracts, err := ctl.stl.ListContracts()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Contracts: %+v", contracts)

	err = c.Render(http.StatusOK, "contracts", ContractsPage{
		Page: Page{
			PageName: "Contracts",
		},
		Contracts: contracts,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
