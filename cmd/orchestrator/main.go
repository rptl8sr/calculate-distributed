package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	app "calculate-distributed/internal/app/orchestrator"
	"calculate-distributed/internal/logger"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	a := app.Must()

	go func() {
		if err := a.Run(ctx); err != nil {
			logger.Error("app.Run: ", "message", err.Error())
			stop()
		}
	}()

	<-ctx.Done()
	logger.Info("Shutdown signal received")

	if err := a.Shutdown(); err != nil {
		logger.Error("app.Close: ", "message", err.Error())
	}

	logger.Info("Shutdown complete")
}
