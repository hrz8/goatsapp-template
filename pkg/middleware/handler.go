package middleware

import (
	"context"
	"strings"

	"github.com/hrz8/goatsapp/internal/port"
	dbrepo "github.com/hrz8/goatsapp/internal/repo/db"
	"github.com/hrz8/goatsapp/views/error_page"
	"github.com/labstack/echo/v4"
)

type projectUsecaser interface {
	GetProjects(ctx context.Context) ([]*dbrepo.GetProjectsRow, error)
}

func EchoContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		return next(&EchoCtx{e})
	}
}

func PopulateActiveProjects(cfg port.AppConfigor, svc projectUsecaser) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			basePath := cfg.GetBasePath()
			path := e.Request().URL.Path

			if basePath != "" {
				basePath = "/" + basePath
			}

			if !strings.HasPrefix(path, basePath+"/assets") {
				projects, err := svc.GetProjects(e.Request().Context())
				if err != nil {
					projects = make([]*dbrepo.GetProjectsRow, 0)
				}
				e.Set("projects", projects)
			}

			return next(e)
		}
	}
}

func NotFound(e echo.Context) error {
	return error_page.NotFoundPage().Render(
		e.Request().Context(),
		e.Response().Writer,
	)
}
