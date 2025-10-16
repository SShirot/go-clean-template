// Package providers provides dependency injection providers for infrastructure
package providers

import (
	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/evrone/go-clean-template/pkg/postgres"
)

// NewLogger creates a new logger instance
func NewLogger(cfg *config.Config) di.LoggerInterface {
	return logger.New(cfg.Log.Level)
}

// NewPostgres creates a new postgres connection
func NewPostgres(cfg *config.Config, l di.LoggerInterface) (di.PostgresInterface, error) {
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(err)
	}
	return pg, nil
}
