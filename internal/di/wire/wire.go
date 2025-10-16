//go:build wireinject
// +build wireinject

// Package wire provides dependency injection with Wire
package wire

import (
	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/internal/di/providers"
	"github.com/google/wire"
)

// InitializeContainer creates a new DI container with all dependencies injected
func InitializeContainer(cfg *config.Config) (*di.Container, error) {
	wire.Build(
		// Infrastructure providers
		providers.NewLogger,
		providers.NewPostgres,

		// Repository providers
		providers.NewTranslationRepo,
		providers.NewUserRepo,

		// Service providers
		providers.NewTranslationService,
		providers.NewUserService,

		// Handler providers
		providers.NewHTTPHandlers,
		providers.NewGRPCHandlers,

		// Container
		di.NewContainer,
	)
	return &di.Container{}, nil
}
