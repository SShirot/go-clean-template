// Package providers provides dependency injection providers for handlers
package providers

import (
	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/internal/interfaces/grpc"
	"github.com/evrone/go-clean-template/internal/interfaces/http"
)

// NewHTTPHandlers creates new HTTP handlers
func NewHTTPHandlers(
	translationService di.TranslationServiceInterface,
	userService di.UserServiceInterface,
) di.HTTPHandlersInterface {
	return http.NewHandlers(translationService, userService)
}

// NewGRPCHandlers creates new gRPC handlers
func NewGRPCHandlers(
	translationService di.TranslationServiceInterface,
	userService di.UserServiceInterface,
) di.GRPCHandlersInterface {
	return grpc.NewHandlers(translationService, userService)
}
