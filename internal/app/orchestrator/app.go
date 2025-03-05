package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"calculate-distributed/internal/api"
	config "calculate-distributed/internal/config/orchestrator"
	"calculate-distributed/internal/logger"
	"calculate-distributed/internal/router"
	"calculate-distributed/internal/server"
	service "calculate-distributed/internal/service/orchestrator"
	"calculate-distributed/internal/storage"
)

type app struct {
	httpServer *http.Server
}

type App interface {
	Run(ctx context.Context) error
	Shutdown() error
}

func Must() App {
	cfg := config.Must()
	logger.Init(cfg.AppLogLevel)
	o := service.New(storage.New(), operationsTimeout(cfg.Timeouts))
	si := server.New(o)
	r := router.New()
	h := api.HandlerFromMux(si, r)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.AppPort),
		Handler: h,
	}

	return &app{
		httpServer: s,
	}
}

func operationsTimeout(timeouts config.Timeouts) map[api.Operation]time.Duration {
	ot := make(map[api.Operation]time.Duration)

	t := reflect.ValueOf(timeouts)
	v := reflect.ValueOf(timeouts)

	for i := 0; i < t.NumField(); i++ {
		ot[api.Operation(t.Type().Field(i).Name)] = v.Field(i).Interface().(time.Duration)
	}

	return ot
}

func (a *app) Run(ctx context.Context) error {
	logger.Info("Starting server", "address", a.httpServer.Addr)

	errChan := make(chan error, 1)

	go func() {
		err := a.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return a.Shutdown()
	}
}

func (a *app) Shutdown() error {
	logger.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		logger.Error("Graceful shutdown failed: ", "error", err.Error())
		return err
	}

	logger.Info("Server and database shutdown completed successfully")
	return nil
}
