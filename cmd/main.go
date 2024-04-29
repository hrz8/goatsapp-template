package main

import (
	"context"
	"log"
	"os"

	"github.com/hrz8/goatsapp/assets"
	"github.com/hrz8/goatsapp/internal/config"
	"github.com/hrz8/goatsapp/internal/port"
	dbrepo "github.com/hrz8/goatsapp/internal/repo/db"
	HomeSvc "github.com/hrz8/goatsapp/internal/service/home"
	ProjectSvc "github.com/hrz8/goatsapp/internal/service/projects"
	SettingSvc "github.com/hrz8/goatsapp/internal/service/settings"
	"github.com/hrz8/goatsapp/views/error_page"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx := context.Background()
	cfg := config.New()
	db := connect(ctx, cfg)

	defer func() {
		db.Close(ctx)
		log.Println("cleaning up...")
		os.Exit(0)
	}()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.RouteNotFound("*", func(c echo.Context) error {
		return error_page.NotFoundPage().Render(
			c.Request().Context(),
			c.Response().Writer,
		)
	})

	base := e.Group(cfg.GetBasePath())
	base.GET("/assets/*", assets.StaticHandler("public"))

	bootstrap(cfg, base, db)

	e.Logger.Fatal(e.Start(":" + cfg.GetAppPort()))
}

func bootstrap(cfg port.AppConfigor, e *echo.Group, db *pgx.Conn) {
	repo := dbrepo.New(db)

	homeUc := HomeSvc.NewUsecase(cfg)
	settingUc := SettingSvc.NewUsecase(cfg)
	projectUc := ProjectSvc.NewUsecase(cfg, repo)

	HomeSvc.NewController(cfg, homeUc).Init(e)
	SettingSvc.NewController(cfg, settingUc).Init(e)
	ProjectSvc.NewController(cfg, projectUc).Init(e)
}
