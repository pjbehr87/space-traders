package stjobs

import (
	"context"

	"github.com/labstack/echo/v4"
	stapi "github.com/pjbehr87/space-traders/st-api"
)

type action interface {
	start(ctx context.Context, ship *stapi.Ship) error
}

type Job struct {
	e *echo.Echo
	ctx *context.Context
	actions   []action
	shipState *stapi.Ship
}

func NewJob(e *echo.Echo, ctx *context.Context, stapi *stapi.APIClient) Job {
	return Job{
		e: e,
		ctx: ctx,
		actions: []action{},
	}
}

func (j Job) run() error {
	shipResp, resp, err := 

	actionLen := len(j.actions)
	for i := 0; ; i = (i + 1) % actionLen {
		err := j.actions[i].start(ctx, ship)
		if err != nil {
			return nil
		}
	}
}
