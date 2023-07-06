package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"strings"
	"time"

	stapi "github.com/pjbehr87/space-traders/st-api"

	"github.com/pjbehr87/space-traders/internal"
	"github.com/pjbehr87/space-traders/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

//go:embed internal/views/*.html
var tmplFS embed.FS

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	funcMap := template.FuncMap{
		"wpHasTrait": func(traits []stapi.WaypointTrait, search string) bool {
			for _, trait := range traits {
				if trait.Symbol == search {
					return true
				}
			}
			return false
		},
		"sysFromWp": func(wp string) string {
			splitWp := strings.Split(wp, "-")
			return fmt.Sprintf("%s-%s", splitWp[0], splitWp[1])
		},
		"timeFmt": func(date time.Time) string {
			localTz, err := time.LoadLocation("Local")
			if err != nil {
				return "bad date: " + err.Error()
			}
			dateLocal := date.In(localTz)
			return dateLocal.Format("2006-01-02 3:4:5 pm")
		},
		"timeUntil": func(date time.Time) string {
			timeUntil := time.Until(date)

			return timeUntil.Round(time.Second).String()
		},
	}
	templates := template.Must(template.New("").Funcs(funcMap).ParseFS(tmplFS, "internal/views/*.html"))
	return &Template{
		templates: templates,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl := template.Must(t.templates.Clone())
	tmpl, err := tmpl.ParseFS(tmplFS, "internal/views/"+name+".html")
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, name+".html", data)
}

//go:embed internal/static/*
var staticFS embed.FS

func main() {
	e := echo.New()

	conf := internal.GetConf()

	// Configure the Logger
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${level}\t")

	logLevel := log.ERROR
	switch strings.ToLower(conf.Server.LogLevel) {
	case "err", "error":
		logLevel = log.ERROR
	case "warn", "warning":
		logLevel = log.WARN
	case "info":
		logLevel = log.INFO
	case "debug":
		logLevel = log.DEBUG

	}
	e.Logger.SetLevel(logLevel)
	e.Logger.Info(fmt.Sprintf("LogLevel set to: %s", conf.Server.LogLevel))

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

	// Log every request at the INFO level. Include query or form params if they are given
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			reqString := "REQUEST uri=%s status=%v"
			values := []interface{}{
				v.URI,
				v.Status,
			}
			paramValues := v.QueryParams
			if len(paramValues) > 0 {
				reqString += " query=%v"
				values = append(values, paramValues)
			}
			formValues := v.FormValues
			if len(formValues) > 0 {
				reqString += " form=%v"
				values = append(values, formValues)
			}

			c.Logger().Info(fmt.Sprintf(reqString, values...))
			return nil
		},
	}))

	// Set the html renderer, that was set up at the top of the page, as the renderer for echo
	t := NewTemplate()
	e.Renderer = t

	// Host embeded the static JS and CSS files
	sfs, err := fs.Sub(staticFS, "internal/static")
	if err != nil {
		panic(err.Error())
	}
	e.Logger.Info("Hosting static files: " + e.StaticFS("/public", sfs).Path)

	// Instantiate the ST-API client
	stConf := stapi.NewConfiguration()
	sta := stapi.NewAPIClient(stConf)

	// Add all of the routs and handlers to the server
	controller.InitRouter(e, sta)

	// Start the server, lisenting at port conf.Server.Port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Server.Port)))
}
