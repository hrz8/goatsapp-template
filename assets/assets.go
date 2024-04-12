package assets

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/hrz8/goatsapp/views/error_page"
	"github.com/labstack/echo/v4"
)

func checkFolder(folder string) error {
	var assetsFolders = []string{"public"}

	folderIsValid := false
	for _, f := range assetsFolders {
		if f == folder {
			folderIsValid = true
		}
	}

	if !folderIsValid {
		return fmt.Errorf("unknown asset folder: " + folder)
	}

	return nil
}

func StaticHandler(folder string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		err = checkFolder(folder)
		if err != nil {
			return err
		}

		path := c.Param("*")
		path, err = url.PathUnescape(path)
		if err != nil {
			return fmt.Errorf("failed to unescape path variable: %w", err)
		}

		if folder == "public" {
			name := "public/" + filepath.ToSlash(filepath.Clean(strings.TrimPrefix(path, "/")))
			err = servePublic(c, name)
		}

		if err != nil && errors.Is(err, echo.ErrNotFound) {
			return error_page.NotFoundPage().Render(c.Request().Context(), c.Response().Writer)
		}

		return err
	}
}
