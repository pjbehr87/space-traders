package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	stapi "github.com/pjbehr87/space-traders/st-api"
	stjobs "github.com/pjbehr87/space-traders/st-jobs"

	"github.com/pjbehr87/space-traders/internal"
	"github.com/pjbehr87/space-traders/internal/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Instantiate the ST-API client
	stConf := stapi.NewConfiguration()
	sta := stapi.NewAPIClient(stConf)

	// Read in the config from the .conf file
	conf := internal.GetConf()

	// Configure log levels and automatic request logging
	internal.LoggerInit(e, conf)

	// Set up HTML renderer and other static code serving
	internal.RendererInit(e)

	// Inject the ST-API auth token into every request context so that the requests can access ST-API
	e.Use(func(fn echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.SetRequest(
				ctx.Request().WithContext(
					context.WithValue(ctx.Request().Context(), stapi.ContextAccessToken, conf.Agent.Token),
				),
			)
			return fn(ctx)
		}
	})

	wg := sync.WaitGroup{}
	exitChan := make(chan bool, 1)
	stj := stjobs.NewRunner(e, sta, conf.Agent.Token, exitChan, &wg)

	// Add all of the routs and handlers to the server
	controller.InitRouter(e, sta, stj)

	// Set up exit listener to close server and job runner
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		e.Logger.Info("Exiting job runner")
		exitChan <- true
		wg.Wait()
		e.Logger.Info("Shutting down server")
		e.Close()
	}()

	// Start the server, lisenting at port conf.Server.Port
	e.Start(fmt.Sprintf(":%d", conf.Server.Port))
}
