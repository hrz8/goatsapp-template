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

func (u *usecase) HandleCreateNewProject(e echo.Context) {
	name := e.FormValue("name")
	alias := e.FormValue("alias")
	description := e.FormValue("description")
	webhookURL := e.FormValue("webhook-url")

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
	u.repo.CreateNewProjects(e.Request().Context(), projects)
	time.Sleep(1 * time.Second)
	// some logic...
}

func (u *usecase) GetAllProjects(ctx context.Context) ([]*dbrepo.Projects, error) {
	return u.repo.GetProjects(ctx)
}
