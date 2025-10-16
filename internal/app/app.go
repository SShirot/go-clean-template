// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/controller/grpc"
	"github.com/evrone/go-clean-template/internal/controller/http"
	"github.com/evrone/go-clean-template/internal/wire"
)

// Run creates objects via constructors using Wire dependency injection.
func Run(cfg *config.Config) {
	// Initialize app with all dependencies using Wire
	app, err := wire.InitializeApp(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to initialize app: %w", err))
	}

	// Ensure postgres connection is closed on exit
	defer app.Postgres.Close()

	// Setup routers
	grpc.NewRouter(app.GRPCServer.App, app.TranslationUC, app.Logger)
	http.NewRouter(app.HTTPServer.App, cfg, app.TranslationUC, app.Logger)

	// Start servers
	app.RabbitMQServer.Start()
	app.NATSServer.Start()
	app.GRPCServer.Start()
	app.HTTPServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		app.Logger.Info("app - Run - signal: %s", s.String())
	case err = <-app.HTTPServer.Notify():
		app.Logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-app.GRPCServer.Notify():
		app.Logger.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	case err = <-app.RabbitMQServer.Notify():
		app.Logger.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	case err = <-app.NATSServer.Notify():
		app.Logger.Error(fmt.Errorf("app - Run - natsServer.Notify: %w", err))
	}

	// Shutdown
	err = app.HTTPServer.Shutdown()
	if err != nil {
		app.Logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = app.GRPCServer.Shutdown()
	if err != nil {
		app.Logger.Error(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
	}

	err = app.RabbitMQServer.Shutdown()
	if err != nil {
		app.Logger.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}

	err = app.NATSServer.Shutdown()
	if err != nil {
		app.Logger.Error(fmt.Errorf("app - Run - natsServer.Shutdown: %w", err))
	}
}
