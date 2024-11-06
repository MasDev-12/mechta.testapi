package dto

import (
	"github.com/google/uuid"
	"time"
)

type URLDto struct {
	Id             uuid.UUID  `json:"id"`
	OriginalURL    string     `json:"original_url"`
	ShortURL       string     `json:"short_url"`
	UserId         uuid.UUID  `json:"user_id"` // Внешний ключ на пользователя
	IsActive       bool       `json:"is_active"`
	ExpiresAt      time.Time  `json:"expires_at"`
	ClickCount     int        `json:"click_count"`
	LastAccessedAt *time.Time `json:"last_accessed_at,omitempty"` // Опускается, если nil
}
