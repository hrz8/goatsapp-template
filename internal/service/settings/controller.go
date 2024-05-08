package settingsvc

import (
	"context"

	"github.com/hrz8/goatsapp/internal/port"
	dbrepo "github.com/hrz8/goatsapp/internal/repo/db"
	"github.com/hrz8/goatsapp/internal/service/settings/page"
	"github.com/hrz8/goatsapp/pkg/core"
	"github.com/labstack/echo/v4"
)

type settingUsecaser interface {
	HandleSettingsPage()
	GetProjects(ctx context.Context) ([]*dbrepo.GetProjectsRow, error)
}

type ctl struct {
	cfg port.AppConfigor
	svc settingUsecaser
}

func NewController(cfg port.AppConfigor, svc settingUsecaser) *ctl {
	return &ctl{cfg, svc}
}

func (c *ctl) Init(g *echo.Group) {
	d := g.Group("/settings")

	// routes
	d.GET("", c.index)
}

func (c *ctl) index(e echo.Context) error {
	ctx := e.Request().Context()
	c.svc.HandleSettingsPage()
	projects, err := c.svc.GetProjects(ctx)
	if err != nil {
		projects = make([]*dbrepo.GetProjectsRow, 0)
	}

	// response as page
	e.Response().Header().Set("HX-Push-Url", core.BaseURL("/settings"))
	return core.RenderView(ctx, e.Response().Writer, page.SettingsPage(projects))
}
