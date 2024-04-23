package port

import "github.com/labstack/echo/v4"

type HomeService interface {
	HomeHandler(c echo.Context) error
}

type SettingService interface {
	SettingsIndex(c echo.Context) error
}
