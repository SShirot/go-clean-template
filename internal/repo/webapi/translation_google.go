package webapi

import (
	"fmt"

	translator "github.com/Conight/go-googletrans"
	"github.com/evrone/go-clean-template/internal/entity"
)

// TranslationWebAPI -.
type TranslationWebAPI struct {
	conf translator.Config
}

// New -.
func New() *TranslationWebAPI {
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"},
		ServiceUrls: []string{"translate.googleapis.com"},
	}

	return &TranslationWebAPI{
		conf: conf,
	}
}

// Translate -.
func (t *TranslationWebAPI) Translate(translation entity.Translation) (entity.Translation, error) {
	trans := translator.New(t.conf)

	result, err := trans.Translate(translation.Original, translation.Source, translation.Destination)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	translation.Translation = result.Text

	return translation, nil
}
