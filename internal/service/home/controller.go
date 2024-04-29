package homesvc

import (
	"github.com/hrz8/goatsapp/internal/port"
	"github.com/hrz8/goatsapp/internal/service/home/page"
	"github.com/hrz8/goatsapp/pkg/core"
	"github.com/labstack/echo/v4"
)

type ctl struct {
	cfg port.AppConfigor
	svc port.HomeHandler
}

func NewController(cfg port.AppConfigor, svc port.HomeHandler) *ctl {
	return &ctl{cfg, svc}
}

func (c *ctl) Init(g *echo.Group) {
	d := g.Group("")

	// routes
	d.GET("", c.index)
}

func (c *ctl) index(e echo.Context) error {
	c.svc.HandleHomePage()

	// response as page
	e.Response().Header().Set("HX-Push-Url", core.BaseURL("/"))
	return core.RenderView(e.Request().Context(), e.Response().Writer, page.HomePage())
}
