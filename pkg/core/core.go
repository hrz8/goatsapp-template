package core

import (
	"context"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func BaseURL(path string) string {
	basePath := os.Getenv("BASE_PATH")
	if basePath != "" {
		return "/" + basePath + path
	}

	return path
}

func RenderView(ctx context.Context, w http.ResponseWriter, component templ.Component) error {
	return component.Render(ctx, w)
}
