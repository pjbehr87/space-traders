package controller

import (
	"net/http"

	stapi "github.com/pjbehr87/space-traders/st-api"

	"github.com/labstack/echo/v4"
)

type contractsController struct {
	sta *stapi.APIClient
}

type contractsPage struct {
	Page PageData

	Contracts []stapi.Contract
}

func NewContractsController(e *echo.Echo, sta *stapi.APIClient) {
	cont := contractsController{
		sta: sta,
	}

	e.Logger.Debug("Router added: Contracts")
	e.GET("/my/contracts", cont.listContracts)
	e.GET("/my/contracts/:contractId", cont.getContract)
	e.POST("/my/contracts/:contractId/accept", cont.acceptContract)
}

func (ctl *contractsController) listContracts(c echo.Context) error {
	myContracts, _, err := ctl.sta.ContractsApi.GetContracts(c.Request().Context()).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Contracts\n%+v", myContracts.Data)

	err = c.Render(http.StatusOK, "myContracts", contractsPage{
		Page: PageData{
			PageName:   "Contracts",
			JavaScript: []string{"contracts"},
		},
		Contracts: myContracts.Data,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *contractsController) getContract(c echo.Context) error {
	contractId := c.Param("contractId")
	contract, _, err := ctl.sta.ContractsApi.GetContract(c.Request().Context(), contractId).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Resp: Contract\n%+v", contract)

	err = c.Render(http.StatusOK, "myContracts", contractsPage{
		Page: PageData{
			PageName:   "Contract",
			JavaScript: []string{"contracts"},
		},
		Contracts: []stapi.Contract{contract.Data},
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}

func (ctl *contractsController) acceptContract(c echo.Context) error {
	contractId := c.Param("contractId")
	_, _, err := ctl.sta.ContractsApi.AcceptContract(c.Request().Context(), contractId).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}

	err = c.NoContent(http.StatusOK)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
