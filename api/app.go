package main

import (
	"api/database"
	"api/database/titlerepo"
	"api/user"
	"context"
	"net/http"

	"api/api"
	"api/config"
	"api/service"
	"go.uber.org/zap"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *api.Server
}

func NewApp(logger *zap.SugaredLogger, settings config.Settings) App {
	pgDb, err := database.NewPgx(settings.Postgres)
	if err != nil {
		panic(err)
	}

	err = database.UpMigrations(pgDb)
	if err != nil {
		panic(err)
	}

	var (
		cli = &http.Client{}

		pingClient = User.NewClient(logger, cli)

		pingRepo = titlerepo.NewRepository(logger, pgDb)

		pingService = service.NewUserService(logger, pingClient, pingRepo)

		server = api.NewServer(logger, settings, pingService)
	)

	return App{
		logger:   logger,
		settings: settings,
		server:   server,
	}
}

func (a App) Run() {
	go func() {
		a.server.Start(a.logger)
	}()
	a.logger.Debugf("HTTP server started on %d", a.settings.Port)
}

func (a App) Stop(ctx context.Context) {
	a.server.Stop(ctx)
	a.logger.Debugf("HTTP server stopped")
}
