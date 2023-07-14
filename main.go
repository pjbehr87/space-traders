package main

import (
	"context"
	"fmt"

	stapi "github.com/pjbehr87/space-traders/st-api"

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

	// Add all of the routs and handlers to the server
	controller.InitRouter(e, sta)

	// Start the server, lisenting at port conf.Server.Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Server.Port)))
}
