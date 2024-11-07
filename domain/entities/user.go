package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID `gorm:"column:id"`
	Username     string    `gorm:"column:user_name, unique;not null"`
	Email        string    `gorm:"column:email, unique;not null"`
	PasswordHash string    `gorm:"column:password_hash, not null"` // Хешированный пароль
	CreatedAt    time.Time `gorm:"column:created_at, not null"`
	IsActive     bool      `gorm:"column:is_active;default:true"`
	URLs         []URL     `gorm:"foreignKey:UserId"` // Связь с URL
}
