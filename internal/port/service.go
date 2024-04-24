package port

import "github.com/labstack/echo/v4"

type HomeService interface {
	HomeIndex(c echo.Context) error
}

type SettingService interface {
	SettingsIndex(c echo.Context) error
}

type AppUsecase interface {
	CreateNewApp(c echo.Context) error
}
