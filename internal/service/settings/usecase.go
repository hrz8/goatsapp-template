package settingsvc

import (
	"context"
	"time"

	"github.com/hrz8/goatsapp/internal/port"
	dbrepo "github.com/hrz8/goatsapp/internal/repo/db"
)

type repositor interface {
	GetProjects(ctx context.Context) ([]*dbrepo.GetProjectsRow, error)
}

type usecase struct {
	cfg  port.AppConfigor
	repo repositor
}

func NewUsecase(cfg port.AppConfigor, repo repositor) *usecase {
	return &usecase{cfg, repo}
}

func (u *usecase) HandleSettingsPage() {
	time.Sleep(1 * time.Second)
	// some logic ...
}

func (u *usecase) GetProjects(ctx context.Context) ([]*dbrepo.GetProjectsRow, error) {
	return u.repo.GetProjects(ctx)
}
