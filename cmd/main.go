package main

import (
	"context"
	"log"
	"os"

	"github.com/hrz8/goatsapp/assets"
	"github.com/hrz8/goatsapp/internal/config"
	dbrepo "github.com/hrz8/goatsapp/internal/repo/db"
	HomeSvc "github.com/hrz8/goatsapp/internal/service/home"
	ProjectSvc "github.com/hrz8/goatsapp/internal/service/projects"
	SettingSvc "github.com/hrz8/goatsapp/internal/service/settings"
	"github.com/hrz8/goatsapp/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
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

	// load services
	repo := dbrepo.New(db)
	homeUc := HomeSvc.NewUsecase(cfg)
	settingUc := SettingSvc.NewUsecase(cfg, repo)
	projectUc := ProjectSvc.NewUsecase(cfg, repo)

	// load default http
	e := echo.New()
	e.Pre(echomw.RemoveTrailingSlash())
	e.Use(echomw.RequestID())
	e.RouteNotFound("*", middleware.NotFound)

	// grouping base path
	base := e.Group(cfg.GetBasePath())

	// load middlewares
	base.Use(middleware.EchoContext)
	base.Use(middleware.PopulateActiveProjects(cfg, projectUc))

	// load web app
	base.GET("/assets/*", assets.StaticHandler("public"))
	HomeSvc.NewController(cfg, homeUc).Init(base)
	SettingSvc.NewController(cfg, settingUc).Init(base)
	ProjectSvc.NewController(cfg, projectUc).Init(base)

	// start http
	e.Logger.Fatal(e.Start(":" + cfg.GetAppPort()))
}
