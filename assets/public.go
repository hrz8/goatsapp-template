package assets

import (
	"embed"
	"errors"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

//go:embed "public/*"
var publicFiles embed.FS

func servePublic(c echo.Context, file string) error {
	f, err := publicFiles.Open(file)
	if err != nil {
		return echo.ErrNotFound
	}

	defer f.Close()

	fi, _ := f.Stat()
	if fi.IsDir() {
		file = filepath.ToSlash(filepath.Join(file, "index.html"))
		f, err = publicFiles.Open(file)
		if err != nil {
			return echo.ErrNotFound
		}
		defer f.Close()
		if fi, err = f.Stat(); err != nil {
			return err
		}
	}

	ff, ok := f.(io.ReadSeeker)
	if !ok {
		return errors.New("file does not implement io.ReadSeeker")
	}

	http.ServeContent(c.Response(), c.Request(), fi.Name(), fi.ModTime(), ff)
	return nil
}
