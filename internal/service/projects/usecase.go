package appsvc

import (
	"context"
	"time"

	"github.com/hrz8/goatsapp/internal/port"
	dbrepo "github.com/hrz8/goatsapp/internal/repo/db"
	"github.com/hrz8/goatsapp/internal/repo/dto"
	"github.com/labstack/echo/v4"
)

type repositor interface {
	GetProjects(ctx context.Context) ([]*dbrepo.Projects, error)
	CreateNewProjects(ctx context.Context, arg []*dbrepo.CreateNewProjectsParams) (int64, error)
}

type usecase struct {
	cfg  port.AppConfigor
	repo repositor
}

func NewUsecase(cfg port.AppConfigor, repo repositor) *usecase {
	return &usecase{cfg, repo}
}

func (u *usecase) HandleCreateNewProject(c echo.Context) {
	name := c.FormValue("name")
	alias := c.FormValue("alias")
	description := c.FormValue("description")
	webhookURL := c.FormValue("webhookURL")

	projects := []*dbrepo.CreateNewProjectsParams{
		{
			Name:        name,
			Alias:       alias,
			Description: &description,
			Settings: dto.ProjectSettings{
				WebhookURL: webhookURL,
			},
		},
	}
	u.repo.CreateNewProjects(c.Request().Context(), projects)
	time.Sleep(1 * time.Second)
	// some logic...
}
