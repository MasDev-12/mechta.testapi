package helpers

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func GenerateShortURL(longURL string) string {
	// Создаем хеш SHA-256
	hash := sha256.Sum256([]byte(longURL))

	// Преобразуем хеш в base64 и обрезаем до нужной длины, например 8 символов
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:8]

	// Удаляем символы, которые могут вызвать проблемы в URL
	shortURL = strings.ReplaceAll(shortURL, "/", "_")
	shortURL = strings.ReplaceAll(shortURL, "+", "-")

	return shortURL
}
