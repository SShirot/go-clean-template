//go:build wireinject
// +build wireinject

// Package wire provides dependency injection for the application
package wire

import (
	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/internal/di/providers"
	"github.com/google/wire"
)

// App is the main application struct
type App struct {
	Logger        di.LoggerInterface
	Postgres      di.PostgresInterface
	HTTPServer    di.HTTPServerInterface
	GRPCServer    di.GRPCServerInterface
	TranslationUC di.TranslationServiceInterface
}

// InitializeApp creates a new App instance with all dependencies injected
func InitializeApp(cfg *config.Config) (*App, error) {
	wire.Build(
		// Infrastructure providers
		providers.NewLogger,
		providers.NewPostgres,
		providers.NewHTTPServer,
		providers.NewGRPCServer,

		// Repository providers
		providers.NewTranslationRepo,

		// Service providers
		providers.NewTranslationService,

		// App struct
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
