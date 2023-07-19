package stjobs

import (
	"context"
	"sync"

	"github.com/labstack/echo/v4"
	stapi "github.com/pjbehr87/space-traders/st-api"
)

type JobRunner struct {
	jobs  chan job
	token string
	sta   *stapi.APIClient
}

func NewJobRunner(e *echo.Echo, sta *stapi.APIClient, token string, exit <-chan bool, wg *sync.WaitGroup) JobRunner {
	jr := JobRunner{
		jobs:  make(chan job),
		token: token,
		sta:   sta,
	}

	jobCtxs := []context.CancelFunc{}
	wg.Add(1)
	go func(jr JobRunner) {
		defer wg.Done()
		for {
			e.Logger.Debug("Waiting for job or exit")
			select {
			case job := <-jr.jobs:
				e.Logger.Info("job added")
				ctx := context.WithValue(context.Background(), stapi.ContextAccessToken, jr.token)
				ctx, ctxCancel := context.WithCancel(ctx)

				jobCtxs = append(jobCtxs, ctxCancel)
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := job.run(e, ctx)
					if err != nil {
						e.Logger.Error(err)
					}
				}()
			case <-exit:
				e.Logger.Info("Canceling all existing contexts")
				for _, cancel := range jobCtxs {
					cancel()
				}
				return
			}
		}
	}(jr)

	return jr
}
