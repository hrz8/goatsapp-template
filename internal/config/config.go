package config

import "os"

type AppConfig interface {
	GetAppPort() string
}

type cfg struct {
	appPort  string
	basePath string
}

func New() *cfg {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "5101"
	}

	basePath := os.Getenv("BASE_PATH")

	return &cfg{
		appPort:  appPort,
		basePath: basePath,
	}
}

func (c *cfg) GetAppPort() string {
	return c.appPort
}

func (c *cfg) GetBasePath() string {
	return c.basePath
}
