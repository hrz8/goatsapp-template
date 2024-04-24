package main

import (
	"log"

	"github.com/hrz8/goatsapp/assets"
	"github.com/hrz8/goatsapp/internal/config"
	AppSvc "github.com/hrz8/goatsapp/internal/service/app"
	HomeSvc "github.com/hrz8/goatsapp/internal/service/home"
	SettingSvc "github.com/hrz8/goatsapp/internal/service/setting"
	"github.com/hrz8/goatsapp/views/error_page"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := config.New()
	appPort := cfg.GetAppPort()
	basePath := cfg.GetBasePath()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.RouteNotFound("*", func(c echo.Context) error {
		return error_page.NotFoundPage().Render(
			c.Request().Context(),
			c.Response().Writer,
		)
	})

	base := e.Group(basePath)
	base.GET("/assets/*", assets.StaticHandler("public"))

	bootstrap(cfg, base)

	e.Logger.Fatal(e.Start(":" + appPort))
}

func bootstrap(cfg config.AppConfig, e *echo.Group) {
	HomeSvc.NewController(cfg, HomeSvc.NewUsecase(cfg)).Init(e)
	SettingSvc.NewController(cfg, SettingSvc.NewUsecase(cfg)).Init(e)
	AppSvc.NewController(cfg, AppSvc.NewUsecase(cfg)).Init(e)
}
