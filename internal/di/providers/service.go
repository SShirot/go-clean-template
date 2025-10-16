// Package providers provides dependency injection providers for services
package providers

import (
	"context"
	"fmt"

	translationApp "github.com/evrone/go-clean-template/internal/application/translation"
	"github.com/evrone/go-clean-template/internal/application/user"
	"github.com/evrone/go-clean-template/internal/di"
	"github.com/evrone/go-clean-template/internal/domain/translation"
	"github.com/evrone/go-clean-template/internal/repo/webapi"
)

// NewTranslationService creates a new translation service
func NewTranslationService(
	repo di.TranslationRepoInterface,
) di.TranslationServiceInterface {
	// Create webapi client
	translateAPI := webapi.New()
	// Convert repo to domain interface
	domainRepo := &domainRepoAdapter{repo: repo}
	return translationApp.NewService(domainRepo, *translateAPI)
}

// domainRepoAdapter adapts DI repo interface to domain interface
type domainRepoAdapter struct {
	repo di.TranslationRepoInterface
}

func (d *domainRepoAdapter) Store(ctx context.Context, translation *translation.Translation) error {
	return d.repo.Store(ctx, translation)
}

func (d *domainRepoAdapter) GetHistory(ctx context.Context, limit, offset int) ([]translation.Translation, error) {
	result, err := d.repo.GetHistory(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	// Convert interface{} to []translation.Translation
	translations := make([]translation.Translation, 0, len(result))
	for _, item := range result {
		if t, ok := item.(translation.Translation); ok {
			translations = append(translations, t)
		}
	}
	return translations, nil
}

func (d *domainRepoAdapter) GetByID(ctx context.Context, id string) (*translation.Translation, error) {
	result, err := d.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if t, ok := result.(*translation.Translation); ok {
		return t, nil
	}
	return nil, fmt.Errorf("invalid translation type")
}

func (d *domainRepoAdapter) Delete(ctx context.Context, id string) error {
	return d.repo.Delete(ctx, id)
}

// NewUserService creates a new user service
func NewUserService(
	repo di.UserRepoInterface,
) di.UserServiceInterface {
	return user.NewService(repo)
}
