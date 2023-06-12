package server

import (
	"path"

	"github.com/labstack/echo/v4"
	swag "github.com/swaggo/echo-swagger"

	"github.com/rytsh/fine/internal/config"
	"github.com/rytsh/fine/internal/server/api"
)

// Register
//
// @title fine API
// @version 1.0
// @description file management service
//
// @contact.name Eray Ates
// @contact.email eates23@gmail.com
//
// @host
// @BasePath /api/v1
func route(e *echo.Echo) {
	basePath := config.App.Server.BasePath
	if basePath != "" {
		basePath = path.Join("/", basePath)
	}

	baseGroup := e.Group(basePath)

	v1 := baseGroup.Group("/api/v1")

	v1.GET("/swagger/*", swag.WrapHandler)

	// set file routes
	api.File(v1)
}
