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
)

func main() {
	cfg := config.New()
	appPort := cfg.GetAppPort()

	e := echo.New()
	e.RouteNotFound("*", func(c echo.Context) error {
		return error_page.NotFoundPage().Render(
			c.Request().Context(),
			c.Response().Writer,
		)
	})

	HomeCtl.New(cfg, HomeSvc.New(cfg)).Init(e)
	SettingsCtl.New(cfg, SettingsSvc.New(cfg)).Init(e)

	e.GET("/assets/*", assets.StaticHandler("public"))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", appPort)))
}
