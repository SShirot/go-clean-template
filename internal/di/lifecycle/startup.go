// Package lifecycle provides application lifecycle management
package lifecycle

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/pkg/logger"
)

// StartupConfig holds startup configuration
type StartupConfig struct {
	Config   *di.Container
	Logger   logger.Interface
	HTTPPort string
	GRPCPort string
}

// Startup handles application startup
func Startup(ctx context.Context, cfg *StartupConfig) error {
	// Initialize infrastructure
	if err := initializeInfrastructure(ctx, cfg); err != nil {
		return fmt.Errorf("failed to initialize infrastructure: %w", err)
	}

	// Start servers
	if err := startServers(ctx, cfg); err != nil {
		return fmt.Errorf("failed to start servers: %w", err)
	}

	cfg.Logger.Info("Application started successfully")
	return nil
}

// initializeInfrastructure sets up database connections, etc.
func initializeInfrastructure(ctx context.Context, cfg *StartupConfig) error {
	// Database migrations, connection health checks, etc.
	cfg.Logger.Info("Infrastructure initialized")
	return nil
}

// startServers starts HTTP, gRPC, and messaging servers
func startServers(ctx context.Context, cfg *StartupConfig) error {
	// Start HTTP server
	cfg.Logger.Info("Starting HTTP server on port %s", cfg.HTTPPort)

	// Start gRPC server
	cfg.Logger.Info("Starting gRPC server on port %s", cfg.GRPCPort)

	// Start messaging servers (RabbitMQ, NATS)
	cfg.Logger.Info("Starting messaging servers")

	return nil
}
