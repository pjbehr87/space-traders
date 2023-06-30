package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"strings"
	"time"

	stlib "github.com/pjbehr87/space-traders/st-lib"

	"github.com/pjbehr87/space-traders/internal"
	"github.com/pjbehr87/space-traders/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

//go:embed internal/views/*.html
var tmplFS embed.FS

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	funcMap := template.FuncMap{
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

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${level}\t")

	conf := internal.GetConf()

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

	t := NewTemplate()
	e.Renderer = t

	e.Logger.Info("Hosting static files: " + e.Static("/public", "internal/static").Path)
	stl := stlib.NewStLib(conf.Agent.Token)

	controller.InitRouter(e, stl)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Server.Port)))
}
