package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type cfg struct {
	appPort     string
	basePath    string
	databaseURL string
}

func New() *cfg {
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "5101"
	}

	basePath := os.Getenv("BASE_PATH")
	databaseURL := os.Getenv("DATABASE_URL")

	return &cfg{
		appPort:     appPort,
		basePath:    basePath,
		databaseURL: databaseURL,
	}
}

func (c *cfg) GetAppPort() string {
	return c.appPort
}

func (c *cfg) GetBasePath() string {
	return c.basePath
}

func (c *cfg) GetDatabaseURL() string {
	return c.databaseURL
}
