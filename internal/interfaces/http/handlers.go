// Package http provides HTTP handlers
package http

import "github.com/evrone/go-clean-template/internal/di"

// handlers implements HTTP handlers
type handlers struct {
	translationService di.TranslationServiceInterface
	userService        di.UserServiceInterface
}

// NewHandlers creates new HTTP handlers
func NewHandlers(translationService di.TranslationServiceInterface, userService di.UserServiceInterface) di.HTTPHandlersInterface {
	return &handlers{
		translationService: translationService,
		userService:        userService,
	}
}

// SetupRoutes sets up HTTP routes
func (h *handlers) SetupRoutes() error {
	// TODO: Implement HTTP route setup
	return nil
}
