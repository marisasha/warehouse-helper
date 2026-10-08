package app

import (
	"context"

	"github.com/marisasha/warehouse-helper/internal/handler"
	"github.com/marisasha/warehouse-helper/internal/repository"
	"github.com/marisasha/warehouse-helper/internal/service"
	httpserver "github.com/marisasha/warehouse-helper/internal/transport/http"
	"gorm.io/gorm"
)

type App struct {
	server   *httpserver.Server
	handlers *handler.Handler
	db       *gorm.DB
}

func NewApp(cfg repository.Config) (*App, error) {
	db, err := repository.NewSQLDB(cfg)
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(httpserver.Server)

	return &App{
		server:   server,
		handlers: handlers,
		db:       db,
	}, nil
}

func (a *App) Run(port string) error {
	return a.server.Run(port, a.handlers.InitRoutes())
}

func (a *App) Shutdown(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}

	sqlDB, err := a.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
