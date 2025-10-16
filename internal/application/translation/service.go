// Package translation defines translation service interface
package translation

import (
	"context"
)

// Service defines the interface for translation business logic
type Service interface {
	// Translate translates text from source to destination language
	Translate(ctx context.Context, req interface{}) (interface{}, error)

	// GetHistory retrieves translation history
	GetHistory(ctx context.Context, limit, offset int) (interface{}, error)

	// GetTranslation retrieves a specific translation
	GetTranslation(ctx context.Context, id string) (interface{}, error)

	// DeleteTranslation removes a translation
	DeleteTranslation(ctx context.Context, id string) error
}
