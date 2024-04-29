package appsvc

import (
	"context"
	"time"

	"github.com/hrz8/goatsapp/internal/port"
	dbrepo "github.com/hrz8/goatsapp/internal/repo/db"
)

type repositor interface {
	GetProjects(ctx context.Context) ([]*dbrepo.Projects, error)
}

type usecase struct {
	cfg  port.AppConfigor
	repo repositor
}

func NewUsecase(cfg port.AppConfigor, repo repositor) *usecase {
	return &usecase{cfg, repo}
}

func (u *usecase) HandleCreateNewProject() {
	time.Sleep(1 * time.Second)
	// some logic...
}
