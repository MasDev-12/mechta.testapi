package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/MasDev-12/mechta.testapi/config"
	"golang.org/x/crypto/argon2"
	"strings"
)

type Argon2Helper struct {
	Settings *config.Argon2Setting
}

func NewArgon2Helper(settings *config.Argon2Setting) *Argon2Helper {
	return &Argon2Helper{
		Settings: settings,
	}
}

// GenerateSalt генерирует случайную соль заданной длины
func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("не удалось сгенерировать соль: %v", err)
	}
	return salt, nil
}

func (a *Argon2Helper) GetPasswordHash(password string) (string, error) {
	salt, err := generateSalt(int(a.Settings.SaltLength))
	if err != nil {
		return "", err
	}
	// Генерация хэша пароля
	passwordHash := argon2.IDKey([]byte(password),
		salt,
		a.Settings.Time,
		a.Settings.Memory,
		a.Settings.Threads,
		a.Settings.KeyLength)
	hashAndSalt := fmt.Sprintf("%s:%s", base64.StdEncoding.EncodeToString(passwordHash), base64.StdEncoding.EncodeToString(salt))
	return hashAndSalt, nil
}
func (a *Argon2Helper) SplitHashAndSalt(hashAndSalt string) (string, string, error) {
	parts := strings.Split(hashAndSalt, ":")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("некорректный формат хэша и соли")
	}
	return parts[0], parts[1], nil
}
