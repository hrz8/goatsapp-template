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
	CustomMiddleware "github.com/hrz8/goatsapp/pkg/middleware"
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

	// load services
	repo := dbrepo.New(db)
	homeUc := HomeSvc.NewUsecase(cfg)
	settingUc := SettingSvc.NewUsecase(cfg)
	projectUc := ProjectSvc.NewUsecase(cfg, repo)

	// load default http
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.RouteNotFound("*", CustomMiddleware.NotFound)

	// grouping base path
	base := e.Group(cfg.GetBasePath())

	// load middlewares
	base.Use(CustomMiddleware.EchoContext)
	base.Use(CustomMiddleware.PopulateActiveProjects(cfg, projectUc))

	// load web app
	base.GET("/assets/*", assets.StaticHandler("public"))
	HomeSvc.NewController(cfg, homeUc).Init(base)
	SettingSvc.NewController(cfg, settingUc).Init(base)
	ProjectSvc.NewController(cfg, projectUc).Init(base)

	// start http
	e.Logger.Fatal(e.Start(":" + cfg.GetAppPort()))
}
