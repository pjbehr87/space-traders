package controller

import (
	"net/http"
	stlib "space-traders/st-lib"

	"github.com/labstack/echo/v4"
)

type AgentController struct {
	stl stlib.StLib
}

func NewAgentController(e *echo.Echo, stl stlib.StLib) {
	cont := AgentController{
		stl: stl,
	}

	e.Logger.Debug("Router added: Agent")
	e.GET("/my/agent", cont.getAgent)
}

func (ctl *AgentController) getAgent(c echo.Context) error {
	c.Logger().Info("GET Agent")
	agent, err := ctl.stl.MyAgent()
	if err != nil {
		c.Logger().Error(err.Error())
		c.String(http.StatusInternalServerError, "Error: "+err.Error())
	}
	c.Logger().Debugf("Agent: %+v", agent)

	err = c.Render(http.StatusOK, "agent", agent)
	if err != nil {
		c.Logger().Error(err.Error())
	}
	return err
}
