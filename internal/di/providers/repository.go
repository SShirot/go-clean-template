// Package providers provides dependency injection providers for repositories
package providers

import (
	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/internal/infrastructure/repository/translation"
	"github.com/evrone/go-clean-template/internal/infrastructure/repository/user"
)

// NewTranslationRepo creates a new translation repository
func NewTranslationRepo(postgres di.PostgresInterface) di.TranslationRepoInterface {
	return translation.NewPostgresRepo(postgres)
}

// NewUserRepo creates a new user repository
func NewUserRepo(postgres di.PostgresInterface) di.UserRepoInterface {
	return user.NewPostgresRepo(postgres)
}
