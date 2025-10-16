// Package di provides dependency injection container
package di

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/config"
)

// Container holds all application dependencies
type Container struct {
	Config *config.Config

	// Infrastructure
	Logger     LoggerInterface
	Postgres   PostgresInterface
	HTTPServer HTTPServerInterface
	GRPCServer GRPCServerInterface

	// Repositories
	TranslationRepo TranslationRepoInterface
	UserRepo        UserRepoInterface

	// Services
	TranslationService TranslationServiceInterface
	UserService        UserServiceInterface

	// Handlers
	HTTPHandlers HTTPHandlersInterface
	GRPCHandlers GRPCHandlersInterface
}

// Interfaces for dependency injection
type LoggerInterface interface {
	Info(msg string, args ...interface{})
	Error(msg interface{}, args ...interface{})
	Fatal(msg interface{}, args ...interface{})
}

type PostgresInterface interface {
	Close()
}

type HTTPServerInterface interface {
	Start()
	Stop() error
	GetApp() interface{}
}

type GRPCServerInterface interface {
	Start()
	Stop() error
	GetServer() interface{}
}

type TranslationRepoInterface interface {
	Store(ctx context.Context, translation interface{}) error
	GetHistory(ctx context.Context, limit, offset int) ([]interface{}, error)
	GetByID(ctx context.Context, id string) (interface{}, error)
	Delete(ctx context.Context, id string) error
}

type UserRepoInterface interface {
	Create(ctx context.Context, user interface{}) error
	GetByID(ctx context.Context, id string) (interface{}, error)
}

type TranslationServiceInterface interface {
	Translate(ctx context.Context, request interface{}) (interface{}, error)
	GetHistory(ctx context.Context, limit, offset int) (interface{}, error)
	GetTranslation(ctx context.Context, id string) (interface{}, error)
	DeleteTranslation(ctx context.Context, id string) error
}

type UserServiceInterface interface {
	CreateUser(ctx context.Context, request interface{}) (interface{}, error)
	GetUser(ctx context.Context, id string) (interface{}, error)
}

type HTTPHandlersInterface interface {
	SetupRoutes() error
}

type GRPCHandlersInterface interface {
	SetupRoutes() error
}

// NewContainer creates a new DI container
func NewContainer(cfg *config.Config) (*Container, error) {
	container := &Container{
		Config: cfg,
	}

	// Initialize dependencies
	if err := container.initializeDependencies(); err != nil {
		return nil, fmt.Errorf("failed to initialize dependencies: %w", err)
	}

	return container, nil
}

// initializeDependencies sets up all dependencies
func (c *Container) initializeDependencies() error {
	// This will be implemented with Wire providers
	return nil
}

// Shutdown gracefully shuts down all dependencies
func (c *Container) Shutdown(ctx context.Context) error {
	// Close database connections
	if c.Postgres != nil {
		c.Postgres.Close()
	}

	return nil
}
