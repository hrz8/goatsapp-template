package homesvc

import (
	"time"

	"github.com/hrz8/goatsapp/internal/port"
)

type usecase struct {
	cfg port.AppConfigor
}

func NewUsecase(cfg port.AppConfigor) *usecase {
	return &usecase{cfg}
}

func (u *usecase) HandleHomePage() {
	time.Sleep(1 * time.Second)
	// some logic...
}
