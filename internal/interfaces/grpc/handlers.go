// Package grpc provides gRPC handlers
package grpc

import "github.com/evrone/go-clean-template/internal/di"

// handlers implements gRPC handlers
type handlers struct {
	translationService di.TranslationServiceInterface
	userService        di.UserServiceInterface
}

// NewHandlers creates new gRPC handlers
func NewHandlers(translationService di.TranslationServiceInterface, userService di.UserServiceInterface) di.GRPCHandlersInterface {
	return &handlers{
		translationService: translationService,
		userService:        userService,
	}
}

// SetupRoutes sets up gRPC routes
func (h *handlers) SetupRoutes() error {
	// TODO: Implement gRPC route setup
	return nil
}
