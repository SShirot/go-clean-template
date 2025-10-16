// Package wire provides dependency injection providers
package wire

import (
	"github.com/evrone/go-clean-template/config"
	amqprpc "github.com/evrone/go-clean-template/internal/controller/amqp_rpc"
	natsrpc "github.com/evrone/go-clean-template/internal/controller/nats_rpc"
	"github.com/evrone/go-clean-template/internal/repo"
	"github.com/evrone/go-clean-template/internal/repo/persistent"
	"github.com/evrone/go-clean-template/internal/repo/webapi"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/internal/usecase/translation"
	"github.com/evrone/go-clean-template/pkg/grpcserver"
	"github.com/evrone/go-clean-template/pkg/httpserver"
	"github.com/evrone/go-clean-template/pkg/logger"
	natsRPCServer "github.com/evrone/go-clean-template/pkg/nats/nats_rpc/server"
	"github.com/evrone/go-clean-template/pkg/postgres"
	rmqRPCServer "github.com/evrone/go-clean-template/pkg/rabbitmq/rmq_rpc/server"
	"github.com/google/wire"
)

// ProviderSet is the complete set of providers for the application
var ProviderSet = wire.NewSet(
	// Infrastructure
	NewLogger,
	NewPostgres,
	NewHTTPServer,
	NewGRPCServer,
	NewRabbitMQServer,
	NewNATSServer,

	// Repositories
	NewTranslationRepo,
	NewTranslationWebAPI,

	// Use Cases
	NewTranslationUseCase,

	// Controllers (routers are set up in app.go)
)

// NewLogger creates a new logger instance
func NewLogger(cfg *config.Config) logger.Interface {
	return logger.New(cfg.Log.Level)
}

// NewPostgres creates a new postgres connection
func NewPostgres(cfg *config.Config, l logger.Interface) (*postgres.Postgres, error) {
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(err)
	}
	return pg, nil
}

// NewHTTPServer creates a new HTTP server
func NewHTTPServer(cfg *config.Config, l logger.Interface) *httpserver.Server {
	return httpserver.New(l, httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
}

// NewGRPCServer creates a new gRPC server
func NewGRPCServer(cfg *config.Config, l logger.Interface) *grpcserver.Server {
	return grpcserver.New(l, grpcserver.Port(cfg.GRPC.Port))
}

// NewRabbitMQServer creates a new RabbitMQ RPC server
func NewRabbitMQServer(cfg *config.Config, translationUseCase usecase.Translation, l logger.Interface) (*rmqRPCServer.Server, error) {
	rmqRouter := amqprpc.NewRouter(translationUseCase, l)
	return rmqRPCServer.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
}

// NewNATSServer creates a new NATS RPC server
func NewNATSServer(cfg *config.Config, translationUseCase usecase.Translation, l logger.Interface) (*natsRPCServer.Server, error) {
	natsRouter := natsrpc.NewRouter(translationUseCase, l)
	return natsRPCServer.New(cfg.NATS.URL, cfg.NATS.ServerExchange, natsRouter, l)
}

// NewTranslationRepo creates a new translation repository
func NewTranslationRepo(pg *postgres.Postgres) repo.TranslationRepo {
	return persistent.New(pg)
}

// NewTranslationWebAPI creates a new translation web API
func NewTranslationWebAPI() repo.TranslationWebAPI {
	return webapi.New()
}

// NewTranslationUseCase creates a new translation use case
func NewTranslationUseCase(repo repo.TranslationRepo, webAPI repo.TranslationWebAPI) usecase.Translation {
	return translation.New(repo, webAPI)
}
