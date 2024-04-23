package controller

import (
	"github.com/hrz8/goatsapp/internal/config"
	"github.com/hrz8/goatsapp/internal/port"
	"github.com/labstack/echo/v4"
)

type ctl struct {
	cfg config.AppConfig
	svc port.SettingService
}

func New(cfg config.AppConfig, svc port.SettingService) *ctl {
	return &ctl{cfg, svc}
}

func (c *ctl) Init(e *echo.Echo) {
	e.GET("/settings", c.svc.SettingsIndex)
}
