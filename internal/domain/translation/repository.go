// Package translation defines translation repository interfaces
package translation

import "context"

// Repository defines the interface for translation data access
type Repository interface {
	// Store stores a new translation
	Store(ctx context.Context, translation *Translation) error

	// GetHistory retrieves translation history
	GetHistory(ctx context.Context, limit, offset int) ([]Translation, error)

	// GetByID retrieves a translation by ID
	GetByID(ctx context.Context, id string) (*Translation, error)

	// Delete removes a translation
	Delete(ctx context.Context, id string) error
}
