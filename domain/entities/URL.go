package entities

import (
	"github.com/google/uuid"
	"time"
)

type URL struct {
	Id             uuid.UUID  `gorm:"column:id"`
	OriginalURL    string     `gorm:"column:origin_url; unique;not null"`
	ShortURL       string     `gorm:"column:short_url; unique;not null"`
	UserId         uuid.UUID  `gorm:"column:user_id; not null"` // Внешний ключ на пользователя
	CreatedAt      time.Time  `gorm:"column:created_at; not null"`
	IsActive       bool       `gorm:"column:is_active;default:true"`
	ExpiresAt      time.Time  `gorm:"column:expires_at; not null"`
	ClickCount     int        `gorm:"column:click_count; default:0"`
	LastAccessedAt *time.Time `gorm:"column:last_accessed_at; default:0"`
}
