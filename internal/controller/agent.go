package controller

import (
	"net/http"

	stapi "github.com/pjbehr87/space-traders/st-api"

	"github.com/labstack/echo/v4"
)

type agentController struct {
	sta *stapi.APIClient
}

type agentPage struct {
	Page PageData

	Agent stapi.Agent
}

func NewAgentController(e *echo.Echo, sta *stapi.APIClient) {
	cont := agentController{
		sta: sta,
	}

	e.Logger.Debug("Router added: Agent\n")
	e.GET("/my/agent", cont.getAgent)
}

func (ctl *agentController) getAgent(c echo.Context) error {
	myAgent, _, err := ctl.sta.AgentsApi.GetMyAgent(c.Request().Context()).Execute()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Resp: Agent\n%+v", myAgent.Data)

	err = c.Render(http.StatusOK, "myAgent", agentPage{
		Page: PageData{
			PageName: "Agent",
		},
		Agent: myAgent.Data,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
