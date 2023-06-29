package controller

import (
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type contractsController struct {
	stl stlib.StLib
}

type contractsPage struct {
	Page PageData

	Contracts []stlib.Contract
}

func NewContractsController(e *echo.Echo, stl stlib.StLib) {
	cont := contractsController{
		stl: stl,
	}

	e.Logger.Debug("Router added: Contracts")
	e.GET("/my/contracts", cont.listContracts)
	e.GET("/my/contracts/:contractId", cont.getContract)
	e.POST("/my/contracts/:contractId/accept", cont.acceptContract)
}

func (ctl *contractsController) listContracts(c echo.Context) error {
	c.Logger().Info("GET MyContracts")
	contracts, err := ctl.stl.ListContracts()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Contracts: %+v", contracts)

	err = c.Render(http.StatusOK, "contracts", contractsPage{
		Page: PageData{
			PageName:   "Contracts",
			JavaScript: []string{"contracts"},
		},
		Contracts: contracts,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *contractsController) getContract(c echo.Context) error {
	c.Logger().Info("GET Contract")
	contract, err := ctl.stl.GetContract(c.Param("contractId"))
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Contract: %+v", contract)

	err = c.Render(http.StatusOK, "contracts", contractsPage{
		Page: PageData{
			PageName:   "Contract",
			JavaScript: []string{"contracts"},
		},
		Contracts: []stlib.Contract{contract},
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *contractsController) acceptContract(c echo.Context) error {
	c.Logger().Info("Post AcceptContract")
	err := ctl.stl.AcceptContract(c.Param("contractId"))
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Accept Contract")

	err = c.NoContent(http.StatusOK)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
