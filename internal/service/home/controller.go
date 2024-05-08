package homesvc

import (
	"github.com/hrz8/goatsapp/internal/port"
	"github.com/hrz8/goatsapp/internal/service/home/page"
	"github.com/hrz8/goatsapp/pkg/core"
	"github.com/labstack/echo/v4"
)

type homeUsecaser interface {
	HandleHomePage()
}

type ctl struct {
	cfg port.AppConfigor
	svc homeUsecaser
}

func NewController(cfg port.AppConfigor, svc homeUsecaser) *ctl {
	return &ctl{cfg, svc}
}

func (c *ctl) Init(g *echo.Group) {
	d := g.Group("")

	// routes
	d.GET("", c.index)
}

func (c *ctl) index(e echo.Context) error {
	ctx := e.Request().Context()
	c.svc.HandleHomePage()

	// response as page
	e.Response().Header().Set("HX-Push-Url", core.BaseURL("/"))
	return core.RenderView(ctx, e.Response().Writer, page.HomePage())
}
