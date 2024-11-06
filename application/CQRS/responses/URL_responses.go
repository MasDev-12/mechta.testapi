package responses

import (
	"github.com/MasDev-12/mechta.testapi/domain/dto"
	"github.com/google/uuid"
	"time"
)

type CreateUrlResponse struct {
	Id       *uuid.UUID `json:"id"`
	ShortURL *string    `json:"short_url"`
	Error    error      `json:"error"`
}

type GetURLResponse struct {
	Id        uuid.UUID `json:"id"`
	ShortURL  string    `json:"short_url"`
	ExpiresAt time.Time `json:"expires_at"`
}

type URLStatsResponse struct {
	OriginalURL    string     `json:"original_url"`
	ClickCount     int        `json:"click_count"`
	LastAccessedAt *time.Time `json:"last_accessed_at,omitempty"` // omitempty для пустого значения
}

type GetUserUrlsResponse struct {
	Urls  []dto.URLDto `json:"urls"`
	Error error        `json:"error"`
}

type GetUrlByShortNameResponse struct {
	Url   *dto.URLDto `json:"url"`
	Error error       `json:"error"`
}

type DeleteUrlByShortNameResponse struct {
	Result bool  `json:"result"`
	Error  error `json:"error"`
}
