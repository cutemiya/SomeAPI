package main

import (
	"api/config"
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	var (
		loggerSugar = logger.Sugar()
		settings    = config.Load(loggerSugar)
	)

	mainCtx := context.Background()
	notifyCtx, cancelFunc := signal.NotifyContext(mainCtx, os.Interrupt, syscall.SIGTERM)
	defer cancelFunc()

	app := NewApp(loggerSugar, settings)
	app.Run()

	select {
	case <-notifyCtx.Done():
	}
	loggerSugar.Debug("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	app.Stop(ctx)
	loggerSugar.Debug("Successful stopped")
}
