package controller

import (
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type agentController struct {
	stl stlib.StLib
}

type agentPage struct {
	Page PageData

	Agent stlib.Agent
}

func NewAgentController(e *echo.Echo, stl stlib.StLib) {
	cont := agentController{
		stl: stl,
	}

	e.Logger.Debug("Router added: Agent")
	e.GET("/my/agent", cont.getAgent)
}

func (ctl *agentController) getAgent(c echo.Context) error {
	c.Logger().Info("GET Agent")
	agent, err := ctl.stl.MyAgent()
	if err != nil {
		c.Logger().Error(err.Error())
		return c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Agent: %+v", agent)

	err = c.Render(http.StatusOK, "agent", agentPage{
		Page: PageData{
			PageName: "Agent",
		},
		Agent: agent,
	})
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
