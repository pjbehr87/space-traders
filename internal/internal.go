package internal

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"math"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	stapi "github.com/pjbehr87/space-traders/st-api"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v3"

	"github.com/labstack/gommon/log"
)

// Go struct representing the config.yaml
type Conf struct {
	Agent struct {
		Token string
		Name  string
	}
	Server struct {
		Port     int
		LogLevel string `yaml:"logLevel"`
	}
}

//go:embed views/*.html
var tmplFS embed.FS

//go:embed static/*
var staticFS embed.FS

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	templates := template.Must(
		template.New("").
			Funcs(TemplateFunctions()).
			ParseFS(tmplFS, "views/*.html"),
	)
	return &Template{
		templates: templates,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl := template.Must(t.templates.Clone())
	tmpl, err := tmpl.ParseFS(tmplFS, "views/"+name+".html")
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, name+".html", data)
}

func TemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"hasContractForGood": func(contracts *[]stapi.Contract, tradeSymbol string) bool {
			if contracts != nil {
				for _, contract := range *contracts {
					for _, deliver := range contract.Terms.Deliver {
						if deliver.TradeSymbol == tradeSymbol {
							return true
						}
					}
				}
			}
			return false
		},
		"ptDist": func(x1 int32, y1 int32, x2 int32, y2 int32) string {
			dist := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
			return fmt.Sprintf("%.2f", dist)
		},
		"i64Commas": func(i int64) string {
			p := message.NewPrinter(language.English)
			return p.Sprintf("%d\n", i)
		},
		"i32Commas": func(i int32) string {
			p := message.NewPrinter(language.English)
			return p.Sprintf("%d\n", i)
		},
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
}

func RendererInit(e *echo.Echo) {
	// Set the html renderer, that was set up at the top of the page, as the renderer for echo
	t := NewTemplate()
	e.Renderer = t

	// Host embeded the static JS and CSS files
	sfs, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err.Error())
	}
	e.Logger.Info("Hosting static files: " + e.StaticFS("/public", sfs).Path)
}

func LoggerInit(e *echo.Echo, conf Conf) {
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
}

// Read the config.yaml into a go object
func GetConf() Conf {
	confFile, err := os.ReadFile("config/conf.yaml")
	if err != nil {
		panic("Read Config File: " + err.Error())
	}

	var c Conf
	if err = yaml.Unmarshal(confFile, &c); err != nil {
		panic("Unmarshal: " + err.Error())
	}

	return c
}
