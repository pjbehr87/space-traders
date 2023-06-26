package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"space-traders/internal"
	stlib "space-traders/st-lib"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

//go:embed internal/views/*.html
var tmplFS embed.FS

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	templates := template.Must(template.New("").ParseFS(tmplFS, "internal/views/*.html"))
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

	t := NewTemplate()
	e.Renderer = t

	e.Logger.SetLevel(log.DEBUG)

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

	stl := stlib.NewStLib(conf.Agent.Token)

	internal.Router(e, stl)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Server.Port)))
}
