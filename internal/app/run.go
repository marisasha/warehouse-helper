package app

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func RunWithGracefulShutdown(application *App, port string) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := application.Run(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("server error: %v", err)
			stop()
		}
	}()

	<-ctx.Done()
	logrus.Info("WAREHOUSE helper API shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := application.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("graceful shutdown failed: %s", err.Error())
	}

	logrus.Info("WAREHOUSE helper API stopped")
}
