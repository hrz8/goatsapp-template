package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hrz8/goatsapp/assets"
	"github.com/hrz8/goatsapp/views/error_page"
	"github.com/hrz8/goatsapp/views/page"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "5101"
	}

	e := echo.New()

	e.RouteNotFound("*", ErrorHandler)
	e.GET("/", HomeHandler)
	e.GET("/settings", SettingsHandler)

	e.GET("/assets/*", assets.StaticHandler("public"))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}

func ErrorHandler(c echo.Context) error {
	return error_page.NotFoundPage().Render(c.Request().Context(), c.Response().Writer)
}

func HomeHandler(c echo.Context) error {
	time.Sleep(1 * time.Second)
	c.Response().Header().Set("HX-Push-Url", "/")
	return page.HomePage().Render(c.Request().Context(), c.Response().Writer)
}

func SettingsHandler(c echo.Context) error {
	time.Sleep(1 * time.Second)
	c.Response().Header().Set("HX-Push-Url", "/settings")
	return page.SettingsPage().Render(c.Request().Context(), c.Response().Writer)
}
