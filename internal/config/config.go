package config

import "os"

type AppConfig interface {
	GetAppPort() string
}

type cfg struct {
	appPort string
}

func New() *cfg {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "5101"
	}

	return &cfg{
		appPort: port,
	}
}

func (c *cfg) GetAppPort() string {
	return c.appPort
}
