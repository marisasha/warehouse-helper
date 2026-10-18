package main

import (
	_ "github.com/marisasha/warehouse-helper/docs"
	"github.com/marisasha/warehouse-helper/internal/app"
	"github.com/marisasha/warehouse-helper/internal/config"

	"github.com/marisasha/warehouse-helper/internal/logger"
	"github.com/sirupsen/logrus"
)

// @title TTL Warehouse helper API
// @version 1.0
// @description API для складского учета
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@ttl-checker.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Введите: Bearer {token}

func main() {
	logger.Init()

	cfg, err := config.Load()
	if err != nil {
		logrus.Fatalf("cannot load config: %s", err)
	}

	application, err := app.NewApp(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize app: %s", err.Error())
	}

	app.RunWithGracefulShutdown(application, cfg.AppPort)

}
