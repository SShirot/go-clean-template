// Package translation defines translation domain entities
package translation

import "time"

// Translation represents a translation entity
type Translation struct {
	ID          string    `json:"id"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Original    string    `json:"original"`
	Translation string    `json:"translation"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TranslationHistory represents a collection of translations
type TranslationHistory struct {
	History []Translation `json:"history"`
	Total   int           `json:"total"`
}

// TranslateRequest represents a translation request
type TranslateRequest struct {
	Source      string `json:"source" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	Original    string `json:"original" validate:"required"`
}

// TranslateResponse represents a translation response
type TranslateResponse struct {
	ID          string `json:"id"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Original    string `json:"original"`
	Translation string `json:"translation"`
}
