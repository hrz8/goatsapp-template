package appsvc

import (
	"fmt"
	"time"

	"github.com/hrz8/goatsapp/internal/port"
	"github.com/hrz8/goatsapp/internal/service/projects/page"
	"github.com/hrz8/goatsapp/pkg/core"
	"github.com/hrz8/goatsapp/views/component"
	"github.com/labstack/echo/v4"
)

type ctl struct {
	cfg port.AppConfigor
	svc port.ProjectHandler
}

func NewController(cfg port.AppConfigor, svc port.ProjectHandler) *ctl {
	return &ctl{cfg, svc}
}

func (c *ctl) Init(g *echo.Group) {
	d := g.Group("/projects")

	// routes
	d.GET("/new", c.newProject)
	d.POST("/store", c.storeApp)
}

func (c *ctl) newProject(e echo.Context) error {
	c.svc.HandleCreateNewProject()

	// response as page
	e.Response().Header().Set("HX-Push-Url", core.BaseURL("/projects/new"))
	return core.RenderView(e.Request().Context(), e.Response().Writer, page.AppNewPage())
}

func (c *ctl) storeApp(e echo.Context) error {
	fmt.Println("lol")
	time.Sleep(1 * time.Second)

	// response
	return core.RenderView(e.Request().Context(), e.Response().Writer, component.Toast(component.ToastProps{
		Message: "successfully",
		Hidden:  false,
	}))
}
