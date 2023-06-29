package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

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
