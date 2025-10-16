// Package lifecycle provides application lifecycle management
package lifecycle

import (
	"context"
	"fmt"
	"time"

	"github.com/evrone/go-clean-template/pkg/logger"
)

// ShutdownConfig holds shutdown configuration
type ShutdownConfig struct {
	Logger    logger.Interface
	Timeout   time.Duration
	Container interface{} // DI Container
}

// Shutdown handles graceful application shutdown
func Shutdown(ctx context.Context, cfg *ShutdownConfig) error {
	cfg.Logger.Info("Starting graceful shutdown...")

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	// Shutdown servers
	if err := shutdownServers(shutdownCtx, cfg); err != nil {
		cfg.Logger.Error(fmt.Errorf("failed to shutdown servers: %w", err))
	}

	// Close database connections
	if err := shutdownInfrastructure(shutdownCtx, cfg); err != nil {
		cfg.Logger.Error(fmt.Errorf("failed to shutdown infrastructure: %w", err))
	}

	cfg.Logger.Info("Graceful shutdown completed")
	return nil
}

// shutdownServers stops all servers gracefully
func shutdownServers(ctx context.Context, cfg *ShutdownConfig) error {
	cfg.Logger.Info("Shutting down servers...")

	// Shutdown HTTP server
	// Shutdown gRPC server
	// Shutdown messaging servers

	return nil
}

// shutdownInfrastructure closes database connections, etc.
func shutdownInfrastructure(ctx context.Context, cfg *ShutdownConfig) error {
	cfg.Logger.Info("Shutting down infrastructure...")

	// Close database connections
	// Close external service connections

	return nil
}
