package appsvc

import (
	"context"

	"github.com/hrz8/goatsapp/internal/port"
	dbrepo "github.com/hrz8/goatsapp/internal/repo/db"
	"github.com/hrz8/goatsapp/internal/repo/dto"
)

type CreateProjectParams struct {
	Name        string
	Alias       string
	Description string
	WebhookURL  string
}

type repositor interface {
	GetProjects(ctx context.Context) ([]*dbrepo.GetProjectsRow, error)
	CreateNewProjects(ctx context.Context, arg []*dbrepo.CreateNewProjectsParams) (int64, error)
}

type usecase struct {
	cfg  port.AppConfigor
	repo repositor
}

func NewUsecase(cfg port.AppConfigor, repo repositor) *usecase {
	return &usecase{cfg, repo}
}

func (u *usecase) HandleCreateNewProject(ctx context.Context, params CreateProjectParams) {
	projects := []*dbrepo.CreateNewProjectsParams{
		{
			Name:        params.Name,
			Alias:       params.Alias,
			Description: &params.Description,
			Settings: dto.ProjectSettings{
				WebhookURL: params.WebhookURL,
			},
		},
	}
	u.repo.CreateNewProjects(ctx, projects)
}

func (u *usecase) GetProjects(ctx context.Context) ([]*dbrepo.GetProjectsRow, error) {
	return u.repo.GetProjects(ctx)
}
