//go:build wireinject
// +build wireinject

// Package wire provides dependency injection
package wire

import (
	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/grpcserver"
	"github.com/evrone/go-clean-template/pkg/httpserver"
	"github.com/evrone/go-clean-template/pkg/logger"
	natsRPCServer "github.com/evrone/go-clean-template/pkg/nats/nats_rpc/server"
	"github.com/evrone/go-clean-template/pkg/postgres"
	rmqRPCServer "github.com/evrone/go-clean-template/pkg/rabbitmq/rmq_rpc/server"
	"github.com/google/wire"
)

// App is the main application struct
type App struct {
	Logger         logger.Interface
	Postgres       *postgres.Postgres
	HTTPServer     *httpserver.Server
	GRPCServer     *grpcserver.Server
	RabbitMQServer *rmqRPCServer.Server
	NATSServer     *natsRPCServer.Server
	TranslationUC  usecase.Translation
}

// InitializeApp creates a new App instance with all dependencies injected
func InitializeApp(cfg *config.Config) (*App, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
