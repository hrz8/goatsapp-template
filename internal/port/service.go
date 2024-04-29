package port

import "github.com/labstack/echo/v4"

type Controller interface {
	Init(g *echo.Group)
}

type HomeHandler interface {
	HandleHomePage()
}

type SettingHandler interface {
	HandleSettingsPage()
}

type ProjectHandler interface {
	HandleCreateNewProject()
}
