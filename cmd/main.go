package main

import (
	"fmt"

	"github.com/hrz8/goatsapp/assets"
	"github.com/hrz8/goatsapp/internal/config"
	HomeCtl "github.com/hrz8/goatsapp/internal/domain/home/controller"
	HomeSvc "github.com/hrz8/goatsapp/internal/domain/home/service"
	SettingsCtl "github.com/hrz8/goatsapp/internal/domain/settings/controller"
	SettingsSvc "github.com/hrz8/goatsapp/internal/domain/settings/service"
	"github.com/hrz8/goatsapp/views/error_page"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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

	HomeCtl.New(cfg, HomeSvc.New(cfg)).Init(base)
	SettingsCtl.New(cfg, SettingsSvc.New(cfg)).Init(base)

	base.GET("/assets/*", assets.StaticHandler("public"))

	fmt.Printf("Server listening on port %s...\n", appPort)
	err := e.Start(":" + appPort)
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}

}
