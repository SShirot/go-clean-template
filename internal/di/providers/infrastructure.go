// Package providers provides dependency injection providers for infrastructure
package providers

import (
	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/pkg/grpcserver"
	"github.com/evrone/go-clean-template/pkg/httpserver"
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

// NewHTTPServer creates a new HTTP server
func NewHTTPServer(cfg *config.Config, l di.LoggerInterface) di.HTTPServerInterface {
	// Cast logger to the expected type
	logger := l.(logger.Interface)
	return &httpServerWrapper{httpserver.New(logger, httpserver.Port(cfg.HTTP.Port))}
}

// NewGRPCServer creates a new gRPC server
func NewGRPCServer(cfg *config.Config, l di.LoggerInterface) di.GRPCServerInterface {
	// Cast logger to the expected type
	logger := l.(logger.Interface)
	return &grpcServerWrapper{grpcserver.New(logger, grpcserver.Port(cfg.GRPC.Port))}
}

// httpServerWrapper wraps httpserver.Server to implement di.HTTPServerInterface
type httpServerWrapper struct {
	*httpserver.Server
}

func (h *httpServerWrapper) Start() {
	h.Server.Start()
}

func (h *httpServerWrapper) GetApp() interface{} {
	return h.Server.App
}

func (h *httpServerWrapper) Stop() error {
	return h.Server.Shutdown()
}

// grpcServerWrapper wraps grpcserver.Server to implement di.GRPCServerInterface
type grpcServerWrapper struct {
	*grpcserver.Server
}

func (g *grpcServerWrapper) Start() {
	g.Server.Start()
}

func (g *grpcServerWrapper) GetServer() interface{} {
	return g.Server.App
}

func (g *grpcServerWrapper) Stop() error {
	return g.Server.Shutdown()
}
