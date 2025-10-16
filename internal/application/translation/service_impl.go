// Package translation implements translation service
package translation

import (
	"context"
	"fmt"
	"time"

	"github.com/evrone/go-clean-template/internal/domain/translation"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/repo/webapi"
	"github.com/google/uuid"
)

// service implements the translation service interface
type service struct {
	repo         translation.Repository
	translateAPI webapi.TranslationWebAPI
}

// NewService creates a new translation service
func NewService(repo translation.Repository, translateAPI webapi.TranslationWebAPI) Service {
	return &service{
		repo:         repo,
		translateAPI: translateAPI,
	}
}

// Translate translates text from source to destination language
func (s *service) Translate(ctx context.Context, req interface{}) (interface{}, error) {
	// Type assert request
	translateReq, ok := req.(*translation.TranslateRequest)
	if !ok {
		return nil, fmt.Errorf("invalid request type")
	}

	// Call external translation API
	translationResult, err := s.translateAPI.Translate(entity.Translation{
		Source:      translateReq.Source,
		Destination: translateReq.Destination,
		Original:    translateReq.Original,
	})
	if err != nil {
		// Fallback: echo original when external API is unavailable
		translationResult = entity.Translation{
			Source:      translateReq.Source,
			Destination: translateReq.Destination,
			Original:    translateReq.Original,
			Translation: translateReq.Original,
		}
	}

	// Create translation entity
	translationEntity := &translation.Translation{
		ID:          uuid.New().String(),
		Source:      translateReq.Source,
		Destination: translateReq.Destination,
		Original:    translateReq.Original,
		Translation: translationResult.Translation,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Store in repository
	if err := s.repo.Store(ctx, translationEntity); err != nil {
		return nil, fmt.Errorf("failed to store translation: %w", err)
	}

	// Return response
	return &translation.TranslateResponse{
		ID:          translationEntity.ID,
		Source:      translationEntity.Source,
		Destination: translationEntity.Destination,
		Original:    translationEntity.Original,
		Translation: translationEntity.Translation,
	}, nil
}

// GetHistory retrieves translation history
func (s *service) GetHistory(ctx context.Context, limit, offset int) (interface{}, error) {
	translations, err := s.repo.GetHistory(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get translation history: %w", err)
	}

	return &translation.TranslationHistory{
		History: translations,
		Total:   len(translations),
	}, nil
}

// GetTranslation retrieves a specific translation
func (s *service) GetTranslation(ctx context.Context, id string) (interface{}, error) {
	translation, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get translation: %w", err)
	}

	return translation, nil
}

// DeleteTranslation removes a translation
func (s *service) DeleteTranslation(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete translation: %w", err)
	}

	return nil
}
