package settingsvc

import (
	"github.com/hrz8/goatsapp/internal/port"
	"github.com/hrz8/goatsapp/internal/service/settings/page"
	"github.com/hrz8/goatsapp/pkg/core"
	"github.com/labstack/echo/v4"
)

type ctl struct {
	cfg port.AppConfigor
	svc port.SettingHandler
}

func NewController(cfg port.AppConfigor, svc port.SettingHandler) *ctl {
	return &ctl{cfg, svc}
}

func (c *ctl) Init(g *echo.Group) {
	d := g.Group("/settings")

	// routes
	d.GET("", c.index)
}

func (c *ctl) index(e echo.Context) error {
	c.svc.HandleSettingsPage()

	// response as page
	e.Response().Header().Set("HX-Push-Url", core.BaseURL("/settings"))
	return core.RenderView(e.Request().Context(), e.Response().Writer, page.SettingsPage())
}
