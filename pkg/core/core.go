package core

import (
	"os"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func BaseURL(path string) string {
	basePath := os.Getenv("BASE_PATH")
	if basePath != "" {
		return "/" + basePath + path
	}

	return path
}

func RenderView(c echo.Context, path string, component templ.Component) error {
	c.Response().Header().Set("HX-Push-Url", BaseURL(path))
	return component.Render(c.Request().Context(), c.Response().Writer)
}
