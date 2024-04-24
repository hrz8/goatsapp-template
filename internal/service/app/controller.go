package appsvc

import (
	"github.com/hrz8/goatsapp/internal/config"
	"github.com/hrz8/goatsapp/internal/port"
	"github.com/labstack/echo/v4"
)

type ctl struct {
	cfg config.AppConfig
	svc port.AppUsecase
}

func NewController(cfg config.AppConfig, svc port.AppUsecase) *ctl {
	return &ctl{cfg, svc}
}

func (c *ctl) Init(g *echo.Group) {
	d := g.Group("/apps")

	// routes
	d.GET("/new", c.svc.CreateNewApp)
}
