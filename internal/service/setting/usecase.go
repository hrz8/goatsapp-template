package settingsvc

import (
	"time"

	"github.com/hrz8/goatsapp/internal/config"
	"github.com/hrz8/goatsapp/internal/service/setting/page"
	"github.com/hrz8/goatsapp/pkg/core"
	"github.com/labstack/echo/v4"
)

type usecase struct {
	cfg config.AppConfig
}

func NewUsecase(cfg config.AppConfig) *usecase {
	return &usecase{cfg}
}

func (u *usecase) SettingsIndex(c echo.Context) error {
	time.Sleep(1 * time.Second)

	return core.RenderView(c, "/settings", page.SettingsPage())
}
