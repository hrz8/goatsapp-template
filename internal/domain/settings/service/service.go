package service

import (
	"time"

	"github.com/hrz8/goatsapp/internal/config"
	"github.com/hrz8/goatsapp/internal/domain/settings/page"
	"github.com/hrz8/goatsapp/pkg/core"
	"github.com/labstack/echo/v4"
)

type svc struct {
	cfg config.AppConfig
}

func New(cfg config.AppConfig) *svc {
	return &svc{cfg}
}

func (u *svc) SettingsIndex(c echo.Context) error {
	time.Sleep(1 * time.Second)

	return core.RenderView(c, "/settings", page.SettingsPage())
}
