package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/rytsh/liz/shutdown"
	"github.com/worldline-go/logz/logecho"
	"github.com/ziflex/lecho/v3"

	"github.com/rytsh/fine/docs"
	"github.com/rytsh/fine/internal/config"
	"github.com/rytsh/fine/internal/fs"
)

var shutdownTimeout = 5 * time.Second

func Start() error {
	docs.SetVersion(config.Name, config.Info.Version, config.App.Server.BasePath)

	// set base path for file system
	fs.BasePath = config.App.Storage.Local.Path

	e := echo.New()

	e.HideBanner = true
	e.Logger = lecho.New(log.Logger)

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(
		middleware.RequestID(),
		middleware.RequestLoggerWithConfig(logecho.RequestLoggerConfig()),
		logecho.ZerologLogger(),
	)

	// add routes
	route(e)

	shutdown.Global.Add("http-server", func() error { return Stop(e) })

	if err := e.Start(fmt.Sprintf("%s:%d", config.App.Server.Host, config.App.Server.Port)); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func Stop(e *echo.Echo) error {
	if e == nil {
		log.Info().Msg("server not running")

		return nil
	}

	log.Info().Msg("stopping service...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := e.Shutdown(ctxShutdown); err != nil {
		return err
	}

	return nil
}
